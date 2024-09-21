package controller

import (
	"QrLanche/backend/pkg/model"
	"QrLanche/backend/pkg/service"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// Handler para criar uma nova mesa
func CreateRestaurantTableHandler(c *gin.Context) {
	var table model.RestaurantTable
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := service.CreateRestaurantTable(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mesa criada com sucesso!",
		"id":      id,
	})
}

// Handler para buscar todas as mesas
func GetAllRestaurantTablesHandler(c *gin.Context) {
	tables, err := service.GetAllRestaurantTables()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mesas listadas com sucesso!",
	})

	for i := 0; i < len(tables); i++ {
		c.JSON(http.StatusOK, gin.H{
			"Table_id":      tables[i].ID,
			"Table_number":  tables[i].Number,
			"Table_available": tables[i].Available,
		})
	}
}

// Handler para buscar uma mesa por ID
func GetRestaurantTableByIDHandler(c *gin.Context) {
	var table model.RestaurantTable

	c.ShouldBindJSON(&table)

	table, err := service.GetRestaurantTableByID(table.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mesa encontrada com sucesso!",
	})

	c.JSON(http.StatusOK, gin.H{
		"Table_id":      table.ID,
		"Table_number":  table.Number,
		"Table_available": table.Available,
	})
}

// Handler para atualizar uma mesa
func UpdateRestaurantTableHandler(c *gin.Context) {
	var table model.RestaurantTable
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.UpdateRestaurantTable(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mesa atualizada com sucesso!",
	})

	c.JSON(http.StatusOK, gin.H{
		"Table_id":      table.ID,
		"Table_number":  table.Number,
		"Table_available": table.Available,
	})
}

// Handler para deletar uma mesa
func DeleteRestaurantTableHandler(c *gin.Context) {
	var table model.RestaurantTable

	err := c.ShouldBindJSON(&table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteRestaurantTable(table.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mesa deletada com sucesso!",
	})

	c.JSON(http.StatusOK, gin.H{
		"Table_id":      table.ID,
		"Table_number":  table.Number,
		"Table_available": table.Available,
	})
}
