package database

import (
	"context"
	"errors"
	"fmt"
	"go-project-frame/dao"
	"go-project-frame/server/config"
	"go-project-frame/server/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	localLog "log"
	"os"
	"time"
)

type InitDBService struct {
	db *gorm.DB
}

func NewInitDBService(db *gorm.DB) *InitDBService {
	return &InitDBService{
		db: db,
	}
}

// InitDB 初始化数据库的入口函数
func (i *InitDBService) InitDB(opts *options.Options) error {
	if len(sysInitSlice) == 0 {
		return errors.New("no initialization procedure is available, please check whether the initialization has completed")
	}
	// 解析config的db信息
	if err := InitDBConfig(opts); err != nil {
		return err
	}
	// 创建mysql的处理对象（TODO 这么写是方便后期实现 多种数据库）
	handler := NewMysqlInitHandler(opts.DB)
	// 初始化db表信息+数据（参数传入了sysInitSlice，包含所有类型的初始化器）
	if err := handler.InitTables(context.TODO(), sysInitSlice); err != nil {
		return err
	}
	return nil
}

// InitDBConfig 初始化数据库配置
func InitDBConfig(opts *options.Options) error {
	// 创建一个Gorm使用的logger对象
	dbLogger := logger.New(
		localLog.New(os.Stdout, "\r\n", localLog.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	// 创建gorm的数据库连接对象
	sqlConfig := config.SysConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.User,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.Name)
	var err error
	if opts.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 关闭自动事务
		SkipDefaultTransaction: false,
		// 进行数据迁移时，是否禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 设置GORM的日志记录器
		Logger: dbLogger,
	}); err != nil {
		return err
	}
	// 获取底层的原生数据库连接对象
	sqlDB, err := opts.DB.DB()
	if err != nil {
		return err
	}
	// 配置连接池参数
	sqlDB.SetMaxIdleConns(config.SysConfig.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.SysConfig.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.SysConfig.Mysql.MaxLifetime) * time.Second)
	// 设置数据库抽象工厂
	opts.Factory = dao.NewShareDaoFactory(opts.DB)
	return nil
}
