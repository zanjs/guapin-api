package model

import (
	"mugg/guapin/app/db"
)

type (
	// Category is
	Category struct {
		IDAutoModel
		TimeAllModel
		Name string `json:"name" gorm:"type:varchar(100);unique"`
		DescriptionModel
		SortModel
		DisabledModel
		Article []*Article `json:"articles"`
	}
	// CategoryArticle is
	CategoryArticle struct {
		CategoryID string `json:"category_id"`
		ArticleID  string `json:"article_id"`
	}
)

// Update is Category
func (m *Category) Update(data *Category) error {
	var (
		err error
	)
	m.Name = data.Name
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (m *Category) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
