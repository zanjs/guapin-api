package model

import (
	"mugg/guapin/app/db"
)

type (
	// Order is
	Order struct {
		IDAutoModel
		GoodsIDModel
		PictureModel
		SortModel
		TimeAllModel
	}
)

// Update is
func (m *Order) Update(data *Order) error {
	var (
		err error
	)
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
func (m *Order) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
