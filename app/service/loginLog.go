package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
)

// LoginLog is
type LoginLog struct {
	mo model.LoginLog
}

// GetAll is
func (s LoginLog) GetAll() ([]model.LoginLog, error) {

	var (
		data []model.LoginLog
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
func (s LoginLog) GetAllQuery(q model.QueryParamsPage) ([]model.LoginLog, error) {

	var (
		data []model.LoginLog
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
func (s LoginLog) GetAllQueryTotal() (int, error) {

	var (
		data  []model.LoginLog
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
func (s LoginLog) GetAllQuerySearch(q model.QueryParamsPage, likeName string) ([]model.LoginLog, error) {

	var (
		data []model.LoginLog
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
func (s LoginLog) GetAllQuerySearchTotal(likeName string) (int, error) {

	var (
		data  []model.LoginLog
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

// Get is LoginLog
func (s LoginLog) Get(id uint64) (model.LoginLog, error) {
	var (
		data model.LoginLog
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

// Create is LoginLog
func (s LoginLog) Create(m *model.LoginLog) error {
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
