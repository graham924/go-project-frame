package user

import (
	"context"
	"go-project-frame/dao/model"
	"gorm.io/gorm"
)

type User interface {
	Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &user{db: db}
}

func (u *user) Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error) {
	return nil, nil
}
