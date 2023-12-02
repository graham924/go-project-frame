package database

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

// DBInitHandler 数据库初始化的顶级接口，也是初始化数据库的入口
type DBInitHandler interface {
	// InitTables 初始化数据库表（建表+初始化数据）
	InitTables(ctx context.Context, initSlice []SystemInitHandler) error
	// createTables 建表（该方法是在InitTables中调用，不对外公开）
	createTables(ctx context.Context, initSlice []SystemInitHandler) error
	// createDatas 初始化数据（该方法是在InitTables中调用，不对外公开）
	createDatas(ctx context.Context, initSlice []SystemInitHandler) error
}

// SystemInitHandler 数据库初始化器 接口
type SystemInitHandler interface {
	// TableCreated 判断所有的表是否都已经完成初始化
	TableCreated(ctx context.Context, db *gorm.DB) bool
	// Migrate 完成所有表的自动迁移（建表）
	Migrate(ctx context.Context, db *gorm.DB) error
	// InitializerData 初始化所有表的数据
	InitializerData(ctx context.Context, db *gorm.DB) error
}

// sysInitSlice 初始化器集合
var sysInitSlice []SystemInitHandler

// RegisterInitHandler 注册初始化器到列表中
func RegisterInitHandler(s SystemInitHandler) {
	if sysInitSlice == nil {
		sysInitSlice = []SystemInitHandler{}
	}
	sysInitSlice = append(sysInitSlice, s)
}
