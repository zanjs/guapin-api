package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"time"
)

// Article is
type Article struct{}

// GetAll is
func (s Article) GetAll() ([]model.Article, error) {

	var (
		data []model.Article
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

// Create is Article
func (s Article) Create(m *model.Article) error {
	var (
		err error
	)
	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Get is Article
func (s Article) Get(id uint64) (model.Article, error) {
	var (
		data model.Article
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

// GetByArticlename is find Article
func (s Article) GetByArticlename(Articlename string) (model.Article, error) {
	var (
		Article model.Article
		err     error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&Article, "name = ?", Articlename).Error; err != nil {
		tx.Rollback()
		return Article, err
	}
	tx.Commit()

	return Article, err
}
