package services

import (
	"errors"
	"mamajulia/internal/database"
	"mamajulia/internal/models"
)

func CreatePedido(pedido models.Order) (models.Order, error) {
	if err := database.DB.Create(&pedido).Error; err != nil {
		return pedido, err
	}

	var pedidoCriado models.Order
	if err := database.DB.Preload("Dishes.Dish").First(&pedidoCriado, pedido.ID).Error; err != nil {
		return pedidoCriado, err
	}

	return pedidoCriado, nil
}

func GetAllPedidos() ([]models.Order, error) {
	var pedidos []models.Order
	if err := database.DB.Preload("Dishes.Dish").Find(&pedidos).Error; err != nil {
		return nil, err
	}
	return pedidos, nil
}

func GetPedidoByID(id uint) (models.Order, error) {
	var pedido models.Order
	if err := database.DB.Preload("Dishes.Dish").First(&pedido, id).Error; err != nil {
		return pedido, errors.New("pedido não encontrado")
	}
	return pedido, nil
}

func UpdatePedidoStatus(id uint, status string) error {
	if status != "em_andamento" && status != "entregue" {
		return errors.New("status inválido. Use 'em_andamento' ou 'entregue'")
	}

	var pedido models.Order
	if err := database.DB.First(&pedido, id).Error; err != nil {
		return errors.New("pedido não encontrado")
	}

	return database.DB.Model(&pedido).Update("status", status).Error
}
