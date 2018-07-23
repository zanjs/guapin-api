package model

type (
	// ShopIDModel is shoop id
	ShopIDModel struct {
		ShopID uint64 `json:"shop_id"`
	}
	// GoodsIDModel is
	GoodsIDModel struct {
		GoodsID uint64 `json:"goods_id"`
	}
	// OrderIDModel is
	OrderIDModel struct {
		OrderID uint64 `json:"order_id"`
	}
	// ExpressNoModel is 物流单号
	ExpressNoModel struct {
		ExpressNo string `json:"express_no"`
	}
	// LogisticsIDModel is 物流商家
	LogisticsIDModel struct {
		LogisticsID uint64 `json:"logistics_id"`
	}
	// GoodsNameModel is 商品名称
	GoodsNameModel struct {
		GoodsName string `json:"goods_name"`
	}
	// GoodsPriceModel is 商品名称
	GoodsPriceModel struct {
		GoodsPrice string `json:"goods_price"`
	}
	// IsGoodsExistsModel is 商品是否有效
	IsGoodsExistsModel struct {
		IsGoodsExists bool `json:"is_goods_exists" gorm:"default:1"`
	}
)
