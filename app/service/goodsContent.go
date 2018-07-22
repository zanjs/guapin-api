package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// GoodsContent is
type GoodsContent struct {
	mo model.GoodsContent
}

// GetAll is
func (s GoodsContent) GetAll() ([]model.GoodsContent, error) {

	var (
		data []model.GoodsContent
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
func (s GoodsContent) GetAllQuerySearch(maps interface{}) ([]model.GoodsContent, error) {

	var (
		data []model.GoodsContent
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
func (s GoodsContent) GetAllQuery(q model.QueryParamsPage) ([]model.GoodsContent, error) {

	var (
		data []model.GoodsContent
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
func (s GoodsContent) GetAllQueryTotal() (int, error) {

	var (
		data  []model.GoodsContent
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

// Get is GoodsContent
func (s GoodsContent) Get(id uint64) (model.GoodsContent, error) {
	var (
		data model.GoodsContent
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

// GetByID is GoodsContent
func (s GoodsContent) GetByID(id uint64) (model.GoodsContent, error) {
	var (
		data model.GoodsContent
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

// GetByGoodsID is GoodsContent
func (s GoodsContent) GetByGoodsID(id uint64) (model.GoodsContent, error) {
	var (
		data model.GoodsContent
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Where("goods_id = ?", id).First(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}

// Create is GoodsContent
func (s GoodsContent) Create(m *model.GoodsContent) error {
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
