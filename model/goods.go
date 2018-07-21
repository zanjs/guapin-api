package model

import (
	"mugg/guapin/app/db"
)

//Goods  type contains page info
type Goods struct {
	IDAutoModel
	CategoryIDModel // 商品分类
	NameModel
	DescriptionModel      // 商品特色描述
	SortModel             // 排序
	RecommendModel        // 推荐状态
	StatusModel           // 商品状态
	PingTuan         bool `json:"pintuan"` // 拼团开启
	PictureModel
	Pictures      []GoodsPicture `json:"pictures"`      // 商品图片
	Content       GoodsContent   `json:"content"`       // 商品详细介绍内容
	OriginalPrice float64        `json:"origina_price"` // 原价
	MinPrice      float64        `json:"min_price"`     // 现价
	PinTuanPrice  float64        `json:"pintuan_price"` // 拼团价
	Stores        uint64         `json:"stores"`        // 库存数
	MinScore      uint64         `json:"min_score"`     // 积分
	Weight        float64        `json:"weight"`        // 重量
	TimeAllModel
	Category GoodsCategory `json:"categories"`
}

// GoodsSeatch is
type GoodsSeatch struct {
	IDAutoModel
	TitleModel
	CategoryIDModel
	DescriptionModel
}

// Update is Goods
func (m *Goods) Update(data *Goods) error {
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
func (m *Goods) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
