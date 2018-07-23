package model

import (
	"mugg/guapin/app/db"
)

// 购物车表 (shoppingcart)
// |-- 自动编号 (id)
// |-- 用户编号 (user_id)
// |-- 商店编号 (shop_id)
// |-- 商品编号 (product_id)
// |-- 是否有效 (is_product_exists)
// |-- 购买数量 (number)
// |-- 创建时间 (created_time)
// 设计说明：商品价格和小计金额是要通过实时关联商品表来读取和计算，因为商户可能会更改商品价格，
// 或者商品已售罄，或者商品已下架等，因此这里只需要记录商品id就可以，商品价格等要实时从商品表读取。

type (
	// ShoppingCart is
	ShoppingCart struct {
		IDAutoModel
		UserIDModel
		ShopIDModel
		GoodsIDModel
		IsGoodsExistsModel
		NumberModel
		TimeAllModel
	}
)

// Update is
func (m *ShoppingCart) Update() error {
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
func (m *ShoppingCart) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
