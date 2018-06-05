package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// Storage is
type Storage struct {
	mo model.Storage
}

// GetAll is
func (s Storage) GetAll() ([]model.Storage, error) {

	var (
		data []model.Storage
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

// GetAllQuery is
func (s Storage) GetAllQuery(q model.QueryParamsPage) ([]model.Storage, error) {

	var (
		data []model.Storage
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
func (s Storage) GetAllQueryTotal() (int, error) {

	var (
		data  []model.Storage
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

// GetAllQuerySearch is
func (s Storage) GetAllQuerySearch(q model.QueryParamsPage, likeName string) ([]model.Storage, error) {

	var (
		data []model.Storage
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Where("name LIKE ?", "%"+likeName+"%").Order("id desc").Offset(q.Offset).Limit(q.Limit).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()
	return data, err
}

// GetAllQuerySearchTotal is
func (s Storage) GetAllQuerySearchTotal(likeName string) (int, error) {

	var (
		data  []model.Storage
		err   error
		count int
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Model(&data).Where("name LIKE ?", "%"+likeName+"%").Count(&count).Error; err != nil {
		tx.Rollback()
		return count, err
	}
	tx.Commit()
	return count, err
}

// Get is Storage
func (s Storage) Get(id uint64) (model.Storage, error) {
	var (
		data model.Storage
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

// Create is Storage
func (s Storage) Create(m *model.Storage) error {
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
