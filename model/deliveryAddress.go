package model

import (
	"mugg/guapin/app/db"
)

// 收货地址表 (delivery_address)
// |-- 自动编号 (address_id)
// |-- 用户编号 (user_id, 用户表自动编号)
// |-- 收件人姓名 (realname)
// |-- 联系电话 (telphone)
// |-- 备用联系电话 (telphone2)
// |-- 国家 (country)
// |-- 省份 (province)
// |-- 城市 (city)
// |-- 地区 (area)
// |-- 街道/详细收货地址 (street)
// |-- 邮政编码 (zip)
// |-- 是否默认收货地址 (is_default_address)
// |-- 创建时间 (created_time)

type (
	// DeliveryAddress is
	DeliveryAddress struct {
		IDAutoModel
		UserIDModel
		Realname  string `json:"realname"`
		Telphone  string `json:"telphone"`
		Telphone2 string `json:"telphone2"`
		Country   string `json:"country"`
		Province  string `json:"province"`
		City      string `json:"city"`
		Area      string `json:"area"`
		Street    string `json:"street"`
		Zip       string `json:"zip"`
		IsDefaultModel
		TimeAllModel
	}
)

// Update is
func (m *DeliveryAddress) Update() error {
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
func (m *DeliveryAddress) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
