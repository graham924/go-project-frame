package model

import (
	"context"
	"gorm.io/gorm"
)

// DBInitializer 数据表初始化器，每一个实例就代表一张表
type DBInitializer interface {
	// TableName 获取表名称
	TableName() string
	// MigrateTable 建表
	MigrateTable(ctx context.Context, db *gorm.DB) error
	// InitializerData 初始化数据
	InitializerData(ctx context.Context, db *gorm.DB) error
	// dataInited 判断数据是否已经完成初始化
	dataInited(ctx context.Context, db *gorm.DB) (bool, error)
}

// InitializerList 所有待初始化表的注册列表
var InitializerList []*OrderedInitializer

// OrderedInitializer 可排序的表信息
type OrderedInitializer struct {
	Order int
	DBInitializer
}

// RegisterInitializer 将表信息，注册到待初始化列表（可以通过指定order，控制表的初始化顺序）
func RegisterInitializer(order int, d DBInitializer) {
	if InitializerList == nil {
		InitializerList = []*OrderedInitializer{}
	}
	InitializerList = append(InitializerList, &OrderedInitializer{
		Order:         order,
		DBInitializer: d,
	})
	InitializerList = bubbleSort(InitializerList)
}

// bubbleSort 对[]*OrderedInitializer，按照order大小进行升序排列
func bubbleSort(s []*OrderedInitializer) []*OrderedInitializer {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j].Order > s[j+1].Order {
				// 使用了多重赋值的特性，将 s[j] 和 s[j+1] 的值进行交换
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
