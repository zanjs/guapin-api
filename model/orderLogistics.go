package model

import (
	"mugg/guapin/app/db"
)

// 订单物流表 (order_logistics)
// |-- 自动编号 (orderlogistics_id)
// |-- 订单编号 (order_id, 订单表自动编号)
// |-- 物流单号 (express_no, 发货快递单号)
// |-- 收货人姓名 (consignee_realname, 收货地址表可能更新或删除，因此要在这里记录)
// |-- 联系电话 (consignee_telphone, 收货地址表可能更新或删除，因此要在这里记录)
// |-- 备用联系电话 (consignee_telphone2, 收货地址表可能更新或删除，因此要在这里记录)
// |-- 收货地址 (consignee_address, 收货地址表可能更新或删除，因此要在这里记录)
// |-- 邮政编码 (consignee_zip, 收货地址表可能更新或删除，因此要在这里记录)
// |-- 物流方式（logistics_type, ems, express）
// |-- 物流商家编号 (logistics_id，物流商家表自动编号)
// |-- 物流发货运费 (logistics_fee，显示给客户的订单运费)
// |-- 快递代收货款费率 (agency_fee, 快递公司代收货款费率，如货值的2%-5%，一般月结)
// |-- 物流成本金额 (delivery_amount, 实际支付给物流公司的金额)
// |-- 物流状态 (orderlogistics_status)
// |-- 物流结算状态 (logistics_settlement_status, 未结算,已结算,部分结算)
// |-- 物流最后状态描述 (logistics_result_last)
// |-- 物流描述 (logistics_result)
// |-- 发货时间 (logistics_create_time)
// |-- 物流更新时间 (logistics_update_time)
// |-- 物流结算时间 (logistics_settlement_time)
// |-- 物流支付渠道
// |-- 物流支付单号
// |-- 物流公司已对账状态 (reconciliation_status，已对账,未对账)
// |-- 物流公司对账日期 (reconciliation_time)
// 设计说明：收货地址可能被修改、删除等，因此这里要记录发货时用户的收货地址，
// 这样就算后来收货地址被删除了，用户在查看历史订单的时候也依然能看到收货地址的快照信息。

type (
	// OrderLogistics is
	OrderLogistics struct {
		IDAutoModel
		OrderIDModel
		ExpressNoModel
		LogisticsIDModel
		ConsigneeCealname  string `json:"consignee_realname"`
		ConsigneeTelphone  string `json:"consignee_telphone"`
		ConsigneeTelphone2 string `json:"consignee_telphone2"`
		ConsigneeAddress   string `json:"consignee_address"`
		ConsigneeZip       string `json:"consignee_zip"`
		TimeAllModel
	}
)

// Update is
func (m *OrderLogistics) Update(data *OrderLogistics) error {
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
func (m *OrderLogistics) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
