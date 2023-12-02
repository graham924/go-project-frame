package database

import (
	"context"
	"go-project-frame/dao/model"
	"gorm.io/gorm"
)

type systemInitHandler struct {
}

func init() {
	// 将系统初始化器，注册到初始化器列表中去
	RegisterInitHandler(&systemInitHandler{})
}

// TableCreated 判断所有的表是否都已经完成初始化
func (s systemInitHandler) TableCreated(ctx context.Context, db *gorm.DB) bool {
	// 获取所有待初始化的表，但凡有一个没有的，就返回false
	for _, initializer := range model.InitializerList {
		if !db.Migrator().HasTable(initializer.TableName()) {
			return false
		}
	}
	// 只有所有待初始化的表都已经初始化，再返回true
	return true
}

// Migrate 完成所有表的自动迁移（建表）
func (s systemInitHandler) Migrate(ctx context.Context, db *gorm.DB) error {
	// 对所有的表，都进行 Migrate迁移操作（建表）
	for _, initializer := range model.InitializerList {
		if err := initializer.MigrateTable(ctx, db); err != nil {
			return err
		}
	}
	return nil
}

// InitializerData 初始化所有表的数据
func (s systemInitHandler) InitializerData(ctx context.Context, db *gorm.DB) error {
	// 对所有的表，都进行 数据初始化操作
	for _, initializer := range model.InitializerList {
		if err := initializer.InitializerData(ctx, db); err != nil {
			return err
		}
	}
	return nil
}
