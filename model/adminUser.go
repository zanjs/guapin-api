package model

import (
	"mugg/guapin/app/db"
)

// Roles is
type Roles struct {
	Roles string `json:"roles"`
}

//AdminUser is table users `gorm:"column:category_id" json:"category_id" form:"category_id"`
type AdminUser struct {
	IDAutoModel
	TimeAllModel
	Name     string `gorm:"column:name;unique_index;default:null" json:"username"`
	Password string `gorm:"default:null" json:"-"` //密码
	// SecretKey string `gorm:"default:null" json:"secret_key"` //密钥
	IsAdmin   bool    `json:"is_admin"`                //是否是管理员
	Avatar    string  `json:"avatar"`                  // 头像链接
	NickName  string  `json:"nick_name"`               // 昵称
	LockState bool    `gorm:"default:'0'" json:"lock"` //锁定状态
	Roles     []Roles `json:"roles"`
}

// AdminUserLogin is
type AdminUserLogin struct {
	Name        string `json:"username"`
	Password    string `json:"password"`     //密码
	OldPassword string `json:"old_password"` //旧密码
}

// AdminUserUpdate is
type AdminUserUpdate struct {
	Name        string `gorm:"column:name" json:"username"`
	Password    string `gorm:"default:null" json:"password"` //密码
	IsAdmin     bool   `json:"is_admin"`                     //是否是管理员
	Avatar      string `json:"avatar"`                       // 头像链接
	NickName    string `json:"nick_name"`                    // 昵称
	LockState   bool   `gorm:"default:'0'" json:"lock"`      //锁定状态
	OldPassword string `json:"old_password"`                 //旧密码
}

// Update is AdminUser
func (m *AdminUser) Update() error {
	var (
		err error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
