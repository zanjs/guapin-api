package model

// Roles is
type Roles struct {
	Roles string `json:"roles"`
}

//User is table users `gorm:"column:category_id" json:"category_id" form:"category_id"`
type User struct {
	IDAutoModel
	TimeAllModel
	Name     string `gorm:"unique_index;default:null" json:"username"`
	Password string `gorm:"default:null" json:"-"` //密码
	// SecretKey string `gorm:"default:null" json:"secret_key"` //密钥
	IsAdmin   bool    `json:"is_admin"`                //是否是管理员
	AvatarURL string  `json:"avatar"`                  // 头像链接
	NickName  string  `json:"nick_name"`               // 昵称
	LockState bool    `gorm:"default:'0'" json:"lock"` //锁定状态
	Roles     []Roles `json:"roles"`
}

// UserLogin is
type UserLogin struct {
	Name        string `json:"username"`
	Password    string `json:"password"`     //密码
	OldPassword string `json:"old_password"` //旧密码
}

// UserLoginLog is
type UserLoginLog struct {
	IDAutoModel
	Name      string `json:"username"`
	UserID    uint64 `json:"user_id"`
	UserAgent string `json:"user_Agent"`
	IPModel
	CreateModel
}
