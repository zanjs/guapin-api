package model

import (
	"mugg/guapin/app/db"
)

//Page type contains page info
type Page struct {
	IDAutoModel
	TitleModel
	SlugModel
	DescriptionModel
	ContentModel
	BodyModel
	ImageModel
	TimeAllModel
}

// Update is Page
func (m *Page) Update(data *Page) error {
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
func (m *Page) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
