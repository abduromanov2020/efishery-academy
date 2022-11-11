package entity

type User struct {
	ID       uint   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Cart     Cart   `gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Cart     Cart   `gorm:"foreignKey:UserID"`
}
