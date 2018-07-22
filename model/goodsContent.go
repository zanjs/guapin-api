package model

import (
	"mugg/guapin/app/db"
)

type (
	// GoodsContent is
	GoodsContent struct {
		IDAutoModel
		GoodsIDModel
		TimeAllModel
		ContentModel
	}
)

// Update is Category
func (m *GoodsContent) Update() error {
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
func (m *GoodsContent) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
