package model

import (
	"mugg/guapin/app/db"
	"time"
)

// 订单表 (order)
// |-- 自动编号（order_id, 自增长主键）
// |-- 订单单号（order_no, 唯一值，供客户查询）
// |-- 商店编号（shop_id, 商店表自动编号）
// |-- 订单状态 (order_status,未付款,已付款,已发货,已签收,退货申请,退货中,已退货,取消交易)
// |-- 商品数量 (product_count, 商品项目数量，不是商品)
// |-- 商品总价 (product_amount_total)
// |-- 订单金额 (order_amount_total，实际付款金额)
// |-- 运费金额 (logistics_fee)
// |-- 是否开箱验货 (is_unpacking_inspection)
// |-- 是否开票（是否开具发票）
// |-- 发票编号 (订单发票表自动编号)
// |-- 收货地址编号 (address_id, 收货地址表自动编号)
// |-- 订单物流编号 (orderlogistics_id, 订单物流表自动编号)
// |-- 订单支付渠道 (pay_channel)
// |-- 订单支付单号 (out_trade_no/escrow_trade_no,第三方支付流水号)
// |-- 创建时间 (下单时间)
// |-- 付款时间
// |-- 发货时间 (ShipTime)
// |-- 签收时间 (SigningTime)
// |-- 客户编号 (user_id，用户表自动编号)
// |-- 客户备注
// |-- 订单结算状态 (order_settlement_status，货到付款、分期付款会用到)
// |-- 订单结算时间 (order_settlement_time)

type (
	// Order is
	Order struct {
		IDAutoModel
		OrderNo string `json:"order_no"`
		UserIDModel
		ShopIDModel
		GoodsIDModel
		GoodsCount       uint    `json:"goods_count"`
		GoodsAmountTotal float64 `json:"goods_amount_total"`
		OrderAmountTotal float64 `json:"order_amount_total"`
		LogisticsFee     float64 `json:"logistics_fee"`
		AddressID        uint64  `json:"address_id"`
		LogisticsID      string  `json:"orderlogistics_id"`
		PayChannel       string  `json:"pay_channel"`
		PayNo            string  `json:"pay_no"`
		RemarkModel
		PayTime     time.Time `json:"pay_at"`
		ShipTime    time.Time `json:"ship_at"`
		SigningTime time.Time `json:"signing_at"`
		TimeAllModel
	}
)

// Update is
func (m *Order) Update() error {
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
func (m *Order) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
