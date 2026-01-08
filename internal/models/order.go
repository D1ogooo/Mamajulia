package models

type Order struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	TableNumber int         `json:"table_number" binding:"required"`
	Status      string      `json:"status" gorm:"default:'em_andamento'"`
	Dishes      []OrderDish `json:"dishes" gorm:"foreignKey:OrderID"`
	CreatedAt   string      `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type OrderDish struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	OrderID  uint `json:"order_id"`
	DishID   uint `json:"dish_id"`
	Quantity int  `json:"quantity" binding:"required,gt=0"`
	Dish     Dish `json:"dish" gorm:"foreignKey:DishID"`
}
