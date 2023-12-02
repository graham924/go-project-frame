package model

import (
	"gorm.io/gorm"
)

// CommonModel 表结构的公共字段（和gorm.Model一致，但json和gorm值更完善）
type CommonModel struct {
	ID int `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	//CreateAt time.Time      `json:"create_at" gorm:"column:create_at"`
	//UpdateAt time.Time      `json:"update_at" gorm:"column:update_at"`
	DeleteAt gorm.DeletedAt `json:"-" gorm:"index"`
}
