package model

import (
	"mugg/guapin/app/db"
)

// LoginLog is
type LoginLog struct {
	IDAutoModel
	Name string `json:"username"`
	UserIDModel
	UserAgent string `json:"user_agent"`
	IPModel
	CreateModel
}

// Update is LoginLog
func (m *LoginLog) Update() error {
	var (
		err error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (m *LoginLog) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
