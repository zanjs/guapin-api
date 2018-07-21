package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// GoodsPicture is
type GoodsPicture struct {
	mo model.GoodsPicture
}

// GetAll is
func (s GoodsPicture) GetAll() ([]model.GoodsPicture, error) {

	var (
		data []model.GoodsPicture
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

// GetAllQuerySearch is
func (s GoodsPicture) GetAllQuerySearch(maps interface{}) ([]model.GoodsPicture, error) {

	var (
		data []model.GoodsPicture
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Where(maps).Order("id desc").Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()
	return data, err
}

// GetAllQuery is
func (s GoodsPicture) GetAllQuery(q model.QueryParamsPage) ([]model.GoodsPicture, error) {

	var (
		data []model.GoodsPicture
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
func (s GoodsPicture) GetAllQueryTotal() (int, error) {

	var (
		data  []model.GoodsPicture
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

// Get is GoodsPicture
func (s GoodsPicture) Get(id uint64) (model.GoodsPicture, error) {
	var (
		data model.GoodsPicture
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

// Create is GoodsPicture
func (s GoodsPicture) Create(m *model.GoodsPicture) error {
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
