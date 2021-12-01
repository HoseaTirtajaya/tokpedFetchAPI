package fetch

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"desc"`
	ImageLink   string  `json:"img_link"`
	Price       float64 `json:"price"`
	ShopName    string  `json:"shop_name"`
}
