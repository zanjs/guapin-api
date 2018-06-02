package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
)

// UserLoginLog is
type UserLoginLog struct{}

// GetAll is
func (s UserLoginLog) GetAll() ([]model.UserLoginLog, error) {

	var (
		data []model.UserLoginLog
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

// Create is UserLoginLog
func (s UserLoginLog) Create(m *model.UserLoginLog) error {
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
