package model

import (
	"mugg/guapin/app/db"
)

// Setting is
type Setting struct {
	IDAutoModel
	Key     string `gorm:"column:key;type:varchar(100);unique" json:"key" form:"key"`
	Name    string `gorm:"column:name" json:"name" form:"name"`
	Value   string `gorm:"column:value" json:"value" form:"value"`
	Details string `gorm:"column:details" json:"details" form:"details"`
	Type    string `gorm:"column:type" json:"type" form:"type"`
	Order   int64  `gorm:"column:order" json:"order" form:"order"`
	Group   string `gorm:"column:group" json:"group" form:"group"`
}

// Update is Setting
func (m *Setting) Update() error {
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
func (m *Setting) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
