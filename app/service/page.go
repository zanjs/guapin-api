package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// Page is
type Page struct {
	mo model.Page
}

// GetAll is
func (s Page) GetAll() ([]model.Page, error) {

	var (
		data []model.Page
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err

}

// Get is Page
func (s Page) Get(id uint64) (model.Page, error) {
	var (
		data model.Page
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Last(&data, id).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// Create is Page
func (s Page) Create(m *model.Page) error {
	var (
		err error
	)
	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
