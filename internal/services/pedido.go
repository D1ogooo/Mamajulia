package controllers

import (
	"mamajulia/internal/models"
	"mamajulia/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePedido(c *gin.Context) {
	var pedido models.Order
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pedido.Status = "em_andamento"

	if err := services.CreatePedido(pedido); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pedido criado com sucesso!", "pedido": pedido})
}

func GetPedidos(c *gin.Context) {
	pedidos, err := services.GetAllPedidos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pedidos)
}

func GetPedidoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	pedido, err := services.GetPedidoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pedido)
}

func UpdatePedidoStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var statusData struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&statusData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdatePedidoStatus(uint(id), statusData.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status do pedido atualizado com sucesso!"})
}
