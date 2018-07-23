package model

import (
	"mugg/guapin/app/db"
)

// 订单商品详情表 (order_detail)
// |-- 自动编号
// |-- 订单编号
// |-- 商品编号
// |-- 商品名称 (product_name, 商品可能删除,所以这里要记录，不能直接读商品表)
// |-- 商品价格 (product_price, 商品可能删除,所以这里要记录)
// |-- 商品型号 (product_marque，前台展示给客户)
// |-- 商品条码 (product_store_barcode, 商品仓库条码)
// |-- 商品型号信息 (product_mode_desc，记录详细商品型号，如颜色、规格、包装等)
// |-- 商品型号参数 (product_mode_params, JSON格式，记录单位编号、颜色编号、规格编号等)
// |-- 折扣比例 (discount_rate 打几折)
// |-- 折扣金额 (discount_amount)
// |-- 购买数量 (number)
// |-- 小计金额 (subtotal)
// |-- 商品是否有效 (is_product_exists)
// |-- 客户商品备注 (remark)
// 设计说明：商品可能被修改、删除等，因此这里要记录下单时用户关注的商品交易摘要信息，
//如价格、数量、型号、型号参数等。这样就算后来商品被删除了，用户在查看历史订单的时候也依然能看到商品的快照信息。

type (
	// OrderDetail is
	OrderDetail struct {
		IDAutoModel
		OrderIDModel
		GoodsIDModel
		GoodsNameModel
		GoodsPriceModel
		IsGoodsExistsModel
		RemarkModel
		Number         uint    `json:"number"`
		DiscountRate   uint    `json:"discount_rate"`
		DiscountAmount float64 `json:"discount_amount"`
		Subtotal       float64 `json:"subtotal"`
		TimeAllModel
	}
)

// Update is
func (m *OrderDetail) Update(data *OrderDetail) error {
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
func (m *OrderDetail) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
