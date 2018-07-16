package model

type (
	// Log is
	Log struct {
		IDAutoModel
		CreateModel
		AppID string `json:"app_id"`
		Name  string `json:"name" gorm:"type:varchar(100);unique"`
		ContentModel
		IPModel
		TypeModel
	}
)
