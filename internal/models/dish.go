package models

type Dish struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Image       string  `json:"image"`
}
