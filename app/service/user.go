package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"mugg/guapin/utils"
	"time"
)

// User is
type User struct{}

// GetAll is
func (s User) GetAll() ([]model.User, error) {

	var (
		data []model.User
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

// Create is user
func (s User) Create(m *model.User) error {
	var (
		err error
	)
	m.Password = utils.HashPassword(m.Password)
	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// GetByUsername is find user
func (s User) GetByUsername(username string) (model.User, error) {
	var (
		user model.User
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&user, "name = ?", username).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, err
}
