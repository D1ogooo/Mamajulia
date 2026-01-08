package models

type Dish struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Value       float64 `json:"value" binding:"required,gt=0"`
	Image       string  `json:"image"`
}
