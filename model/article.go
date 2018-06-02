package model

import (
	"mugg/guapin/app/db"
)

//Article  type contains page info
type Article struct {
	IDAutoModel
	TitleModel
	CategoryIDModel
	DescriptionModel
	PictureModel
	ContentModel
	TimeAllModel
}

// Update is Article
func (m *Article) Update(data *Article) error {
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

// Delete is
func (m *Article) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
