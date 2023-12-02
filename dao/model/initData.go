package model

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"go-project-frame/pkg/consts"
)

// 数据库表的初始化顺序
const (
	SysUserOrder = iota
)

var (
	// SysUserEntities sys_user表的mock数据
	SysUserEntities = []*SysUser{
		{
			UUID:        uuid.NewV4(),
			UserName:    "admin",
			Password:    "$2a$14$Zfb6w0UDBFMN0.nJeVXCUO3zH/iWKGtbBYyIzDDRnC..EgTS0Et0S",
			NickName:    "admin",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: uint(consts.SystemUserAuthorityIdAdmin),
			Phone:       "12345678901",
			Email:       "grahamzhu@tencent.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
		{
			UUID:        uuid.NewV4(),
			UserName:    "chenteng",
			Password:    "$2a$14$yLCxKYP46M2NRnXujYe3mOfNe00GtBtjpaLM2eIzYCzYKQXqzsuka",
			NickName:    "chenteng",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: uint(consts.SystemUserAuthorityIdDefault),
			Phone:       "12345678901",
			Email:       "test@qq.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
		{
			UUID:        uuid.NewV4(),
			UserName:    "chentengsub",
			Password:    "$2a$14$MPINiht5QO2wlR3DynizXOtuqcNAOrNZdrSUKXrbjqcKbK.jcfyAW",
			NickName:    "chentengsub",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: uint(consts.SystemUserAuthorityIdSubDefault),
			Phone:       "12345678901",
			Email:       "test@qq.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
	}
)
