package dao

import (
	"go-project-frame/dao/user"
	"gorm.io/gorm"
)

type ShareDaoFactory interface {
	GetDB() *gorm.DB
	User() user.User
}

// NewShareDaoFactory 创建一个共享的数据库工厂对象
func NewShareDaoFactory(db *gorm.DB) ShareDaoFactory {
	return &shareDaoFactory{
		db: db,
	}
}

type shareDaoFactory struct {
	db *gorm.DB
}

func (s *shareDaoFactory) GetDB() *gorm.DB {
	return s.db
}

func (s *shareDaoFactory) User() user.User {
	return user.NewUser(s.db)
}
