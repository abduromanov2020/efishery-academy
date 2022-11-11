package entity

type Category struct {
	ID           int        `gorm:"primaryKey;column:id" json:"id"`
	CategoryName string     `gorm:"column:category_name; not null" json:"category_name"`
	Products     []Products `gorm:"foreignKey:CategoryID"`
}

type CreateCategoryRequest struct {
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}

type UpdateCategoryRequest struct {
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}

type GetCategoryResponse struct {
	ID           int    `gorm:"primaryKey;column:id" json:"id"`
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}
