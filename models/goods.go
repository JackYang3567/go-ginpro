package models
 
type Goods struct {
	GoodsId	int64  `gorm:"column:goodsId",json:"goodsid"` // 商品id
	GoodsName string  `gorm:"column:goodsName",json:"goodsname"` //商品名字
	Subject string `gorm:"column:subject",json:"subject"`     //商品描述
	Price string `gorm:"column:price",json:"price"`           //商品价格
	Stock int `gorm:"column:stock",json:"stock"`              //商品库存
}
 
func (Goods) TableName() string {
	return "goods"

}