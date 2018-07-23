package service

import (
	"mugg/guapin/app/db"
	"mugg/guapin/model"
	"mugg/guapin/utils"
	"time"
)

// AdminUser is
type AdminUser struct{}

// GetAll is
func (s AdminUser) GetAll() ([]model.AdminUser, error) {

	var (
		data []model.AdminUser
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
func (s AdminUser) Create(m *model.AdminUser) error {
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

// GetAllQuery is
func (s AdminUser) GetAllQuery(q model.QueryParamsPage) ([]model.AdminUser, error) {

	var (
		data []model.AdminUser
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Offset(q.Offset).Limit(q.Limit).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()
	return data, err
}

// GetAllQueryTotal is
func (s AdminUser) GetAllQueryTotal() (int, error) {

	var (
		data  []model.AdminUser
		err   error
		count int
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Model(&data).Count(&count).Error; err != nil {
		tx.Rollback()
		return count, err
	}
	tx.Commit()
	return count, err
}

// GetAllQuerySearch is
func (s AdminUser) GetAllQuerySearch(q model.QueryParamsPage, likeName string) ([]model.AdminUser, error) {

	var (
		data []model.AdminUser
		err  error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Where("name LIKE ?", "%"+likeName+"%").Order("id desc").Offset(q.Offset).Limit(q.Limit).Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()
	return data, err
}

// GetAllQuerySearchTotal is
func (s AdminUser) GetAllQuerySearchTotal(likeName string) (int, error) {

	var (
		data  []model.AdminUser
		err   error
		count int
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Model(&data).Where("name LIKE ?", "%"+likeName+"%").Count(&count).Error; err != nil {
		tx.Rollback()
		return count, err
	}
	tx.Commit()
	return count, err
}

// GetByAdminUsername is find user
func (s AdminUser) GetByAdminUsername(username string) (model.AdminUser, error) {
	var (
		user model.AdminUser
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
