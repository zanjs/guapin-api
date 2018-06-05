package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
)

// Setting is
type Setting struct {
	mo model.Setting
}

// GetAll is
func (s Setting) GetAll() ([]model.Setting, error) {

	var (
		data []model.Setting
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
func (s Setting) GetAllQuery(q model.QueryParamsPage) ([]model.Setting, error) {

	var (
		data []model.Setting
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
func (s Setting) GetAllQueryTotal() (int, error) {

	var (
		data  []model.Setting
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

// Get is Setting
func (s Setting) Get(id uint64) (model.Setting, error) {
	var (
		data model.Setting
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

// Create is Setting
func (s Setting) Create(m *model.Setting) error {
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
