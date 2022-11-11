package entity

type Products struct {
	ID         int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Image      string `json:"image"`
	CategoryID int    `json:"category_id"`
	Category   Category
}

type CreateProductsRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Image      string `json:"image"`
	CategoryID int    `json:"category_id"`
}

type UpdateProductsRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Image      string `json:"image"`
	CategoryID int    `json:"category_id"`
}

type GetProductsResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Image    string `json:"image"`
	Category Category
}
