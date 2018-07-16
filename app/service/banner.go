package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// Banner is
type Banner struct{}

// GetAll is
func (s Banner) GetAll() ([]model.Banner, error) {

	var (
		data []model.Banner
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
func (s Banner) GetAllQuery(q model.QueryParamsPage) ([]model.Banner, error) {

	var (
		data []model.Banner
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

// GetAllQuerySearch is
func (s Banner) GetAllQuerySearch(q model.QueryParamsPage, maps interface{}, likeTitle string) ([]model.Banner, error) {

	var (
		data []model.Banner
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Where(maps).Where("name LIKE ?", "%"+likeTitle+"%").Order("id desc").Offset(q.Offset).Limit(q.Limit).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()
	return data, err
}

// GetAllQueryTotal is
func (s Banner) GetAllQueryTotal() (int, error) {

	var (
		data  []model.Banner
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

// GetAllQuerySearchTotal is
func (s Banner) GetAllQuerySearchTotal(maps interface{}, likeTitle string) (int, error) {

	var (
		data  []model.Banner
		err   error
		count int
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Model(&data).Where(maps).Where("name LIKE ?", "%"+likeTitle+"%").Count(&count).Error; err != nil {
		tx.Rollback()
		return count, err
	}
	tx.Commit()
	return count, err
}

// Create is Banner
func (s Banner) Create(m *model.Banner) error {
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

// Get is Banner
func (s Banner) Get(id uint64) (model.Banner, error) {
	var (
		data model.Banner
		err  error
		// category model.Category
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Preload("Category").Last(&data, id).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// GetID is Banner
func (s Banner) GetID(id uint64) (model.Banner, error) {
	var (
		data model.Banner
		err  error
		// category model.Category
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&data, id).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// GetByBannername is find Banner
func (s Banner) GetByBannername(Bannername string) (model.Banner, error) {
	var (
		Banner model.Banner
		err    error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&Banner, "name = ?", Bannername).Error; err != nil {
		tx.Rollback()
		return Banner, err
	}
	tx.Commit()

	return Banner, err
}
