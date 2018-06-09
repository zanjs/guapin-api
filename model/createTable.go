package model

import (
	"mugg/guapin/app/db"
)

//CreateTable is init db table
func CreateTable() error {
	gorm.MysqlConn().AutoMigrate(&User{},
		&Page{},
		&Category{},
		&Article{},
		&Storage{},
		&LoginLog{},
		&Setting{},
	)
	return nil
}
