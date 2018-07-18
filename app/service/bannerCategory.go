package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// BannerCategory is
type BannerCategory struct {
	mo model.BannerCategory
}

// GetAll is
func (s BannerCategory) GetAll() ([]model.BannerCategory, error) {

	var (
		data []model.BannerCategory
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
func (s BannerCategory) GetAllQuerySearch(maps interface{}) ([]model.BannerCategory, error) {

	var (
		data []model.BannerCategory
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
func (s BannerCategory) GetAllQuery(q model.QueryParamsPage) ([]model.BannerCategory, error) {

	var (
		data []model.BannerCategory
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
func (s BannerCategory) GetAllQueryTotal() (int, error) {

	var (
		data  []model.BannerCategory
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

// Get is BannerCategory
func (s BannerCategory) Get(id uint64) (model.BannerCategory, error) {
	var (
		data model.BannerCategory
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

// Create is BannerCategory
func (s BannerCategory) Create(m *model.BannerCategory) error {
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
