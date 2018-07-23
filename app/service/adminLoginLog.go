package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
)

// AdminLoginLog is
type AdminLoginLog struct {
	mo model.AdminLoginLog
}

// GetAll is
func (s AdminLoginLog) GetAll() ([]model.AdminLoginLog, error) {

	var (
		data []model.AdminLoginLog
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
func (s AdminLoginLog) GetAllQuery(q model.QueryParamsPage) ([]model.AdminLoginLog, error) {

	var (
		data []model.AdminLoginLog
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
func (s AdminLoginLog) GetAllQueryTotal() (int, error) {

	var (
		data  []model.AdminLoginLog
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
func (s AdminLoginLog) GetAllQuerySearch(q model.QueryParamsPage, likeName string) ([]model.AdminLoginLog, error) {

	var (
		data []model.AdminLoginLog
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
func (s AdminLoginLog) GetAllQuerySearchTotal(likeName string) (int, error) {

	var (
		data  []model.AdminLoginLog
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

// Get is AdminLoginLog
func (s AdminLoginLog) Get(id uint64) (model.AdminLoginLog, error) {
	var (
		data model.AdminLoginLog
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

// Create is AdminLoginLog
func (s AdminLoginLog) Create(m *model.AdminLoginLog) error {
	var (
		err error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
