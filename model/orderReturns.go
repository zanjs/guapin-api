package model

import (
	"mugg/guapin/app/db"
)

// 订单退货表 (order_returns)
// |-- 自动编号 (order_returns_id)
// |-- 退货编号 (returns_no，供客户查询)
// |-- 订单编号 (order_id, 订单表自动编号)
// |-- 物流单号 (express_no, 退货物流单号)
// |-- 收货人姓名 (consignee_realname)
// |-- 联系电话 (consignee_telphone)
// |-- 备用联系电话 (consignee_telphone2)
// |-- 收货地址 (consignee_address)
// |-- 邮政编码 (consignee_zip)
// |-- 物流方式（logistics_type, ems, express）
// |-- 物流商家编号
// |-- 物流发货运费 (logistics_fee，退货运费)
// |-- 物流状态 (orderlogistics_status)
// |-- 物流最后状态描述
// |-- 物流描述
// |-- 物流更新时间
// |-- 物流发货时间
// |-- 退货类型 (returns_type, 全部退单,部分退单)
// |-- 退货处理方式 (handling_way, PUPAWAY:退货入库;REDELIVERY:
//	重新发货;RECLAIM-REDELIVERY:不要求归还并重新发货; REFUND:退款; COMPENSATION:不退货并赔偿)
// |-- 退款金额 (returns_amount)
// |-- 退货销售员承担的费用 (seller_punish_fee)
// |-- 退货申请时间 (return_submit_time)
// |-- 退货处理时间 (handling_time)
// |-- 退货原因
// 设计说明：退货可能被修改、删除等，因此这里要记录退货时商家的退货地址信息，

type (
	// OrderReturns is
	OrderReturns struct {
		IDAutoModel
		OrderReturnsNo string `json:"order_no"`
		GoodsIDModel
		PictureModel
		SortModel
		TimeAllModel
	}
)

// Update is
func (m *OrderReturns) Update(data *OrderReturns) error {
	var (
		err error
	)
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
func (m *OrderReturns) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
