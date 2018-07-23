package model

import "time"

type (
	// IDAutoModel is
	IDAutoModel struct {
		ID uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"unique_index;not null;unique;primary_key;column:id"`
	}
	// PIDAutoModel is
	PIDAutoModel struct {
		PID uint64 `json:"pid"`
	}
	// RecommendModel is 推荐状态
	RecommendModel struct {
		RecommendStatus int `json:"recommend_status"`
	}
	// RemarkModel is 备注
	RemarkModel struct {
		Remark string `json:"remark"`
	}
	// IsExistsModel is 是否有效
	IsExistsModel struct {
		IsExists bool `json:"is_exists" gorm:"default:1"`
	}
	// IsDefaultModel is 是否默认
	IsDefaultModel struct {
		IsDefault bool `json:"is_default"`
	}
	// NumberModel is 数量
	NumberModel struct {
		Number uint `json:"number"`
	}
	// LevelModel is
	LevelModel struct {
		Level uint64 `json:"level"`
	}
	// IDModel is
	IDModel struct {
		ID string `json:"id" sql:"index"  gorm:"unique_index;not null;unique;primary_key;column:id"`
	}
	// UUIDModel is
	UUIDModel struct {
		UID string `json:"uid" sql:"index"  gorm:"unique_index;not null;unique;primary_key;column:uid"`
	}
	// UserIDModel 用户ID
	UserIDModel struct {
		UserID uint64 `json:"user_id"`
	}
	AdminUserIDModel struct {
		AdminUserID uint64 `json:"admin_user_id"`
	}
	// CreateModel is
	CreateModel struct {
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	}
	// UpdatedAtModel is
	UpdatedAtModel struct {
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	}
	// DeletedAtModel is
	DeletedAtModel struct {
		DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
	}
	// TimeAllModel is
	TimeAllModel struct {
		CreateModel
		UpdatedAtModel
		DeletedAtModel
	}

	// IPModel is
	IPModel struct {
		IP string `json:"ip"`
	}
	// PictureModel is 图片
	PictureModel struct {
		Picture string `json:"picture"`
	}
	// DescriptionModel is
	DescriptionModel struct {
		Description string `json:"description"`
	}
	// TitleModel is
	TitleModel struct {
		Title string `json:"title"`
	}

	// ContentModel is
	ContentModel struct {
		Content string `json:"content" gorm:"type:longtext"`
	}

	// BodyModel is
	BodyModel struct {
		Body string `json:"body" gorm:"type:longtext"`
	}
	// NameModel is
	NameModel struct {
		Name string `json:"name"`
	}
	// CategoryIDModel is
	CategoryIDModel struct {
		CategoryID uint64 `json:"category_id"`
	}
	// SlugModel is
	SlugModel struct {
		Slug string `json:"slug" gorm:"type:varchar(100);unique"`
	}
	// ImageModel is
	ImageModel struct {
		Image string `json:"image"`
	}
	// StatusModel is
	StatusModel struct {
		Status int `json:"status"`
	}
	// StatuStrModel is
	StatuStrModel struct {
		StatuStr string `json:"status_str"`
	}
	// SortModel is
	SortModel struct {
		Sort int `json:"sort"`
	}
	// MarkModel is
	MarkModel struct {
		Mark string `json:"mark"`
	}
	// TypeModel is
	TypeModel struct {
		Type string `json:"type"`
	}
	// ParentIDModel is
	ParentIDModel struct {
		ParentID int64 `json:"parent_id"`
	}
	// SizeModel is
	SizeModel struct {
		Size int64 `json:"size"`
	}
	// SizeStrModel is
	SizeStrModel struct {
		SizeStr string `json:"size_str"`
	}
	// URLModel is
	URLModel struct {
		URL string `json:"url"`
	}
	// LinkModel is
	LinkModel struct {
		Link string `json:"link"`
	}
	// DisabledModel is
	DisabledModel struct {
		Disabled bool `json:"disabled" gorm:"default:'0'"`
	}
	// BusinessIDModel is 业务编号
	BusinessIDModel struct {
		BusinessID string `json:"business_id"`
	}
	// TypeMKModel is 类型标记
	TypeMKModel struct {
		TypeMK string `json:"type"`
	}
	// QueryParams is
	QueryParams struct {
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
		WareroomID int    `json:"wareroom_id"`
		Day        int    `json:"day"`
		ProductID  int    `json:"product_id"`
		Count      int    `json:"count"`
		Num        int    `json:"num"`
	}
	// QueryParamsPage is
	QueryParamsPage struct {
		Page   int `json:"page"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)
