package database

import (
	"context"
	"gorm.io/gorm"
)

// mysqlInitHandler mysql初始化处理器
type mysqlInitHandler struct {
	db *gorm.DB
}

// NewMysqlInitHandler new a mysqlInitHandler
func NewMysqlInitHandler(db *gorm.DB) *mysqlInitHandler {
	return &mysqlInitHandler{
		db: db,
	}
}

// InitTables 初始化数据库表（建表+初始化数据）
func (m *mysqlInitHandler) InitTables(ctx context.Context, initSlice []SystemInitHandler) error {
	// createTables 建表
	if err := m.createTables(ctx, initSlice); err != nil {
		return err
	}
	// createDatas 初始化数据
	if err := m.createDatas(ctx, initSlice); err != nil {
		return err
	}
	return nil
}

// createTables 建表（该方法是在InitTables中调用，不对外公开）
func (m *mysqlInitHandler) createTables(ctx context.Context, initSlice []SystemInitHandler) error {
	// 遍历所有的初始化器
	for _, initHandler := range initSlice {
		// 如果这个表已经被创建了，则跳过
		if initHandler.TableCreated(ctx, m.db) {
			continue
		}
		// 如果表没有被创建，则进行表迁移（创建或更新表）
		if err := initHandler.Migrate(ctx, m.db); err != nil {
			return err
		}
	}
	return nil
}

// createDatas 初始化数据（该方法是在InitTables中调用，不对外公开）
func (m *mysqlInitHandler) createDatas(ctx context.Context, initSlice []SystemInitHandler) error {
	// 遍历所有的初始化器
	for _, initHandler := range initSlice {
		// 使用初始化器，对数据进行初始化
		if err := initHandler.InitializerData(ctx, m.db); err != nil {
			return err
		}
	}
	return nil
}
