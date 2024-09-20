package controller

import (
	"QrLanche/backend/pkg/model"
	"QrLanche/backend/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrderHandler(c *gin.Context){
	var order model.Order

	if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	id, err := service.CreateOrder(order)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o item: " + err.Error()})
        return
    }

	// Busca os nomes e preços dos itens usando os IDs
	itens, err := service.GetNamesAndPricesByIds(order.Itens)
	if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
        "message": "Erro ao buscar itens do pedido",
        "error":   err.Error(),
    })
    return
}

	c.JSON(http.StatusOK, gin.H{
        "message": "order criado com sucesso!",
        "orderID": id,
		"itens":itens,
    })
}

func DeleteOrderHandler(c *gin.Context){
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	err:= service.DeleteOrder(order.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar order",
			"error":   err.Error(),
		})
		return
	}

		c.JSON(http.StatusOK, gin.H{
			"message": "order deletado com sucesso!",
		})
}

func GetAllOrdersHandler(c *gin.Context){
	orders, err := service.GetAllOrders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Listagem completa!",
    })
	
	
    for i := 0; i < len(orders); i++ {
		items, err := service.GetNamesAndPricesByIds(orders[i].Itens)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

        c.JSON(http.StatusOK, gin.H{
            "order_id":orders[i].ID,
            "order_customer":orders[i].CustomerID,
            "order_itens":items,
            "order_table":orders[i].TableID,
            "order_total":orders[i].Total,
        })
    }
}

func SelectOrderByIdHandler(c *gin.Context){
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	selectedOrder, err := service.SelectOrderById(order.ID)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	items, err := service.GetNamesAndPricesByIds(selectedOrder.Itens)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

        c.JSON(http.StatusOK, gin.H{
            "order_id":selectedOrder.ID,
            "order_customer":selectedOrder.CustomerID,
            "order_itens":items,
            "order_table":selectedOrder.TableID,
            "order_total":selectedOrder.Total,
        })
}

func UpdateOrderHandler(c *gin.Context){
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    err := service.UpdateOrder(order)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    selectedOrder, err := service.SelectOrderById(order.ID)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Item atualizado com sucesso!",
    })

    c.JSON(http.StatusOK, gin.H{
        "new_order": selectedOrder,
    })
}