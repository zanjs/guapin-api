package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// GoodsCategory is
type GoodsCategory struct {
	mo model.GoodsCategory
}

// GetAll is
func (s GoodsCategory) GetAll() ([]model.GoodsCategory, error) {

	var (
		data []model.GoodsCategory
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
func (s GoodsCategory) GetAllQuerySearch(maps interface{}) ([]model.GoodsCategory, error) {

	var (
		data []model.GoodsCategory
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Where(maps).Order("sort desc,id desc").Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()
	return data, err
}

// GetAllQuery is
func (s GoodsCategory) GetAllQuery(q model.QueryParamsPage) ([]model.GoodsCategory, error) {

	var (
		data []model.GoodsCategory
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
func (s GoodsCategory) GetAllQueryTotal() (int, error) {

	var (
		data  []model.GoodsCategory
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

// Get is GoodsCategory
func (s GoodsCategory) Get(id uint64) (model.GoodsCategory, error) {
	var (
		data model.GoodsCategory
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

// Create is GoodsCategory
func (s GoodsCategory) Create(m *model.GoodsCategory) error {
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
