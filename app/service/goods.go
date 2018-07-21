package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// Goods is
type Goods struct{}

// GetAll is
func (s Goods) GetAll() ([]model.Goods, error) {

	var (
		data []model.Goods
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
func (s Goods) GetAllQuery(q model.QueryParamsPage) ([]model.Goods, error) {

	var (
		data []model.Goods
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
func (s Goods) GetAllQuerySearch(q model.QueryParamsPage, maps interface{}, likeTitle string) ([]model.Goods, error) {

	var (
		data []model.Goods
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
func (s Goods) GetAllQueryTotal() (int, error) {

	var (
		data  []model.Goods
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
func (s Goods) GetAllQuerySearchTotal(maps interface{}, likeTitle string) (int, error) {

	var (
		data  []model.Goods
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

// Create is Goods
func (s Goods) Create(m *model.Goods) error {
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

// Get is Goods
func (s Goods) Get(id uint64) (model.Goods, error) {
	var (
		data model.Goods
		err  error
		// category model.Category
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Preload("Category").Preload("Pictures").Preload("Content").Last(&data, id).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// GetID is Goods
func (s Goods) GetID(id uint64) (model.Goods, error) {
	var (
		data model.Goods
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

// GetByGoodsname is find Goods
func (s Goods) GetByGoodsname(Goodsname string) (model.Goods, error) {
	var (
		Goods model.Goods
		err   error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&Goods, "name = ?", Goodsname).Error; err != nil {
		tx.Rollback()
		return Goods, err
	}
	tx.Commit()

	return Goods, err
}
