package cartModel

type AddProductToCart struct {
	CartId    string `json:"cart_id" form:"cart_id"`
	ProductId string `json:"product_id" form:"product_id"`
	Quantity  int64  `json:"quantity" form:"quantity"`
	OutletId  string `json:"outlet_id" form:"outlet_id" `
}
