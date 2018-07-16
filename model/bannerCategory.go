package model

import (
	"mugg/guapin/app/db"
)

type (
	// BannerCategory is
	BannerCategory struct {
		IDAutoModel
		TimeAllModel
		Name string `json:"name" gorm:"type:varchar(100);unique"`
		DescriptionModel
		SortModel
		DisabledModel
		TypeModel
		Banner []*Banner `json:"banners,omitempty"`
	}
)

// Update is Category
func (m *BannerCategory) Update(data *BannerCategory) error {
	var (
		err error
	)
	m.Name = data.Name
	m.Description = data.Description
	m.Disabled = data.Disabled
	m.Type = data.Type
	m.Sort = data.Sort
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (m *BannerCategory) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
