package model

import (
	"mugg/guapin/app/db"
)

type (
	// GoodsCategory is
	GoodsCategory struct {
		IDAutoModel
		TimeAllModel
		Name string `json:"name" gorm:"type:varchar(100);"`
		PIDAutoModel
		TypeModel
		LevelModel
		DescriptionModel
		SortModel
		DisabledModel
		List []*GoodsCategory `json:"list,omitempty"`
	}
)

// Update is
func (m *GoodsCategory) Update(data *GoodsCategory) error {
	var (
		err error
	)
	m.Name = data.Name
	m.Description = data.Description
	m.Disabled = data.Disabled
	m.Sort = data.Sort
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (m *GoodsCategory) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
