package controllers

import (
	"mamajulia/internal/models"
	"mamajulia/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPratos(c *gin.Context) {
	pratos, err := services.GetAllPratos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pratos)
}

func GetPratoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	prato, err := services.GetPratoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, prato)
}

func CreatePratos(c *gin.Context) {
	var prato models.Dish
	if err := c.ShouldBindJSON(&prato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if prato.Image == "" {
		prato.Image = "https://via.placeholder.com/400x300?text=" + prato.Name
	}

	if err := services.CreatePrato(prato); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Prato criado com sucesso!", "prato": prato})
}

func UpdatePratos(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var prato models.Dish
	if err := c.ShouldBindJSON(&prato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdatePrato(uint(id), prato); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prato atualizado com sucesso!"})
}

func DeletePratos(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := services.DeletePrato(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prato deletado com sucesso!"})
}
