package model

import (
	"context"
	"database/sql"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	CommonModel
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	UserName    string    `json:"userName" gorm:"index;comment:用户登陆名"`
	Password    string    `json:"-" gorm:"comment:用户登陆密码"` // 转换为JSON时忽略该字段
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`
	Avatar      string    `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`
	AuthorityId uint      `json:"authorityId" gorm:"default:2222;comment:用户角色ID"`
	// Authority 这里原本是有个外键，被我去掉了
	Authority string `json:"authority" gorm:"not null;comment:用户角色"`
	// Authorities 这里原本是个many2many，被我去掉了
	Authorities string        `json:"authorities" gorm:"comment:用户该角色同时具有哪些角色的权限"`
	Phone       string        `json:"phone" gorm:"comment:用户手机号"`
	Email       string        `json:"email" gorm:"comment:用户邮箱"`
	Enable      int           `json:"enable" gorm:"default:1;comment:账号可用性 0冻结 1正常"`
	Status      sql.NullInt64 `json:"status" gorm:"type:int(11);comment:用户在线状态 0离线 1在线"`
}

func init() {
	RegisterInitializer(SysUserOrder, &SysUser{})
}

// TableName 获取表名称
func (u *SysUser) TableName() string {
	return "sys_user"
}

// MigrateTable 建表
func (u *SysUser) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&u)
}

// InitializerData 初始化数据
func (u *SysUser) InitializerData(ctx context.Context, db *gorm.DB) error {
	if ok, err := u.dataInited(ctx, db); ok || err != nil {
		return err
	}
	if err := db.WithContext(ctx).Create(SysUserEntities).Error; err != nil {
		return err
	}
	//if err := db.Model(&SysUserEntities[0]).Association("Authorities").Replace(SysAuthorityEntities); err != nil {
	//	return err
	//}
	//if err := db.Model(&SysUserEntities[1]).Association("Authorities").Replace(SysAuthorityEntities[:1]); err != nil {
	//	return err
	//}
	return nil
}

// dataInited 判断数据是否已经完成初始化
func (u *SysUser) dataInited(ctx context.Context, db *gorm.DB) (bool, error) {
	// 判断admin用户是否已经完成初始化，是的话，就算已经完成初始化了
	adminOk, err := u.adminInited(ctx, db)
	if err != nil {
		return false, nil
	}
	return adminOk, nil
}

func (u *SysUser) adminInited(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *SysUser
	if err := db.WithContext(ctx).Where("user_name = 'admin'").Find(&out).Error; err != nil {
		return false, err
	}
	return out.ID != 0, nil
}
