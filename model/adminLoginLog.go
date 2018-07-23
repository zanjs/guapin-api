package model

import (
	"mugg/guapin/app/db"
)

// AdminLoginLog is
type AdminLoginLog struct {
	IDAutoModel
	Name string `json:"username"`
	AdminUserIDModel
	AdminUserAgent string `json:"user_agent"`
	IPModel
	CreateModel
}

// Update is AdminLoginLog
func (m *AdminLoginLog) Update() error {
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
func (m *AdminLoginLog) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
