package web

import "time"

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Owner       string    `json:"owner"`
	Discount    int       `json:"discount"`
	Description string    `json:"description"`
	ExpireAt    time.Time `json:"expire_at"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
}

type OrderResponse struct {
	UUID        string    `json:"uuid"`
	Username    string    `json:"user_name"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Owner       string    `json:"owner"`
	Discount    int       `json:"discount"`
	Description string    `json:"description"`
	ExpireAt    time.Time `json:"expire_at"`
	Status      string    `json:"status"`
}

type CategoryResponse struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Product     []ProductResponse `json:"product"`
}

type CategoryCreated struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserRegister struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLogin struct {
	Token string `json:"token"`
}

type Pagination struct {
	Metadata Metadata    `json:"_metadata"`
	Data     interface{} `json:"data"`
}

type Metadata struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	TotalData int64 `json:"data_count"`
	TotalPage int64 `json:"page_count"`
}

type PaginationElastic struct {
	Metadata MetadataElastic `json:"_metadata"`
	Data     interface{}     `json:"data"`
}

type MetadataElastic struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type PaymentResponse struct {
	OrderUUID          string    `json:"uuid"`
	ProductName        string    `json:"product_name"`
	ProductPrice       int       `json:"product_price"`
	ProductOwner       string    `json:"product_owner"`
	ProductExpire      time.Time `json:"product_expire"`
	ProductDiscount    int       `json:"product_discount"`
	ProductDescription string    `json:"product_description"`
	TotalHarga         int       `json:"total_harga"`
	Status             string    `json:"status"`
}

type BuyResponse struct {
	ProductName        string `json:"product_name"`
	ProductPrice       int    `json:"product_price"`
	ProductDiscount    int    `json:"product_discount"`
	ProductDescription string `json:"product_description"`
	TotalHarga         int    `json:"total_harga"`
	PayURL             string `json:"pay_url"`
}

type Cart struct {
	ProductName        string    `json:"product_name"`
	ProductPrice       int       `json:"product_price"`
	ProductOwner       string    `json:"product_owner"`
	ProductExpire      time.Time `json:"product_expire"`
	ProductDiscount    int       `json:"product_discount"`
	ProductDescription string    `json:"product_description"`
	TotalBeli          int       `json:"total_beli"`
	URLCheckout        string    `json:"url_checkout"`
}
