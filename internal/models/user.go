package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"-" binding:"required,min=6"`
	Role     string `json:"role" gorm:"default:'user'"`
}
