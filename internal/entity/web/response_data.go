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
}

type CategoryResponse struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Product     []ProductResponse `json:"product"`
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
