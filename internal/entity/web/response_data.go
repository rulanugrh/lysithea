package web

import "time"

type ProductCreated struct {
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

type GetProduct struct {
	Product ProductCreated `json:"product"`
}

type GetAllProduct struct {
	Product []ProductCreated `json:"list_product"`
}

type OrderCreated struct {
	UUID        string    `json:"uuid"`
	Username    string    `json:"user_name"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Owner       string    `json:"owner"`
	Discount    int       `json:"discount"`
	Description string    `json:"description"`
	ExpireAt    time.Time `json:"expire_at"`
}

type GetOrder struct {
	Order OrderCreated `json:"order"`
}

type GetAllOrder struct {
	Order []OrderCreated `json:"list_order"`
}
