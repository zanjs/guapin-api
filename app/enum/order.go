package enum

// OrderStatus is 订单状态枚举
type OrderStatus int

const (
	//PayNot is 未支付
	PayNot OrderStatus = 1 + iota
	//PayOK is 支付成功 已经付款
	PayOK
	//Ship is 已发货
	Ship
	// Signing is 签收
	Signing
	// Cancel is 取消交易
	Cancel
)

// LogisticsType is 配送物流方式
type LogisticsType int

const (
	// EMS is
	EMS LogisticsType = 1 + iota
	// Express is 快递
	Express
	// Oneself is 上门自取
	Oneself
)
