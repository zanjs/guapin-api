package model

import (
	"mugg/guapin/app/db"
)

type (
	// Banner is
	Banner struct {
		IDAutoModel
		TypeModel
		BusinessIDModel
		CategoryIDModel
		DescriptionModel
		NameModel
		PictureModel
		LinkModel
		SortModel
		StatusModel
		StatuStrModel
		TimeAllModel
		Category BannerCategory `json:"categories"`
	}
)

// Update is Banner
func (m *Banner) Update(data *Banner) error {
	var (
		err error
	)
	// m.CategoryID = data.CategoryID
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (m *Banner) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
