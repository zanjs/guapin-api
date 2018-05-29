package service

import (
	"mugg/guapin/app/db"
)

// Base is
type Base struct{}

// All is
func (b Base) All(mo interface{}) (interface{}, error) {
	var (
		data interface{}
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&mo).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err
}
