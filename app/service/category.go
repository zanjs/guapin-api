package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// Category is
type Category struct {
	mo model.Category
}

// GetAll is
func (s Category) GetAll() ([]model.Category, error) {

	var (
		data []model.Category
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err

}

// GetAllQuery is
func (s Category) GetAllQuery(q model.QueryParamsPage) ([]model.Category, error) {

	var (
		data []model.Category
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Offset(q.Offset).Limit(q.Limit).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()
	return data, err
}

// GetAllQueryTotal is
func (s Category) GetAllQueryTotal() (int, error) {

	var (
		data  []model.Category
		err   error
		count int
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Model(&data).Count(&count).Error; err != nil {
		tx.Rollback()
		return count, err
	}
	tx.Commit()
	return count, err
}

// Get is Category
func (s Category) Get(id uint64) (model.Category, error) {
	var (
		data model.Category
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

// Create is Category
func (s Category) Create(m *model.Category) error {
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
