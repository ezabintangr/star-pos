package cartModel

type AddProductToCart struct {
	ProductId string `json:"product_id" form:"product_id"`
	Quantity  int64  `json:"quantity" form:"quantity"`
	OutletId  string `json:"outlet_id" form:"outlet_id" `
}
