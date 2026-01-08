package models

import "time"

type Order struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	TableNumber int         `json:"table_number" binding:"required"`
	Status      string      `json:"status" gorm:"default:'em_andamento'"`
	Dishes      []OrderDish `json:"dishes" gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderDish struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	OrderID  uint `json:"order_id"`
	DishID   uint `json:"dish_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required,gt=0"`
	Dish     Dish `json:"dish" gorm:"foreignKey:DishID"`
}
