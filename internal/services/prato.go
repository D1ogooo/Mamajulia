package services

import (
	"errors"
	"mamajulia/internal/database"
	"mamajulia/internal/models"
)

func GetAllPratos() ([]models.Dish, error) {
	var pratos []models.Dish
	if err := database.DB.Find(&pratos).Error; err != nil {
		return nil, err
	}
	return pratos, nil
}

func GetPratoByID(id uint) (models.Dish, error) {
	var prato models.Dish
	if err := database.DB.First(&prato, id).Error; err != nil {
		return prato, errors.New("prato não encontrado")
	}
	return prato, nil
}

func CreatePrato(prato models.Dish) error {
	return database.DB.Create(&prato).Error
}

func UpdatePrato(id uint, prato models.Dish) error {
	var existingPrato models.Dish
	if err := database.DB.First(&existingPrato, id).Error; err != nil {
		return errors.New("prato não encontrado")
	}

	return database.DB.Model(&existingPrato).Updates(prato).Error
}

func DeletePrato(id uint) error {
	var prato models.Dish
	if err := database.DB.First(&prato, id).Error; err != nil {
		return errors.New("prato não encontrado")
	}
	return database.DB.Delete(&prato).Error
}
