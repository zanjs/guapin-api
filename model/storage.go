package model

import (
	"mugg/guapin/app/db"
)

// Storage is
type Storage struct {
	IDAutoModel
	UUIDModel
	NameModel
	TypeModel
	URLModel
	UserID int64 `json:"user_id"`
	SizeModel
	SizeStrModel
	TimeAllModel
}

// Update is Storage
func (m *Storage) Update() error {
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
func (m *Storage) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
