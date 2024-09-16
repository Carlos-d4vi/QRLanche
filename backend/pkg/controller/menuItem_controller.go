package controller

import (
	"QrLanche/backend/pkg/model"
	"QrLanche/backend/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateMenuItemHandler(c *gin.Context) {
    var item model.MenuItem

    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, err := service.CreateMenuItem(item)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o item: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Pedido criado com sucesso!",
        "orderID": id,
    })
}

func GetAllMenuItemsHandler(c *gin.Context){
    items, err := service.GetAllMenuItems()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Listagem completa!",
    })

    for i := 0; i < len(items); i++ {
        c.JSON(http.StatusOK, gin.H{
            "item_id":items[i].ID,
            "item_name":items[i].Name,
            "item_price":items[i].Price,
        })
    }
}

func SelectItemByIdHandler(c *gin.Context){
    var item model.MenuItem

    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
        return
    }

    selectedItem, err := service.SelectItemById(item)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao selecionar o item: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Item encontrado!",
    })

    c.JSON(http.StatusOK, gin.H{
        "item_name":selectedItem.Name,
        "item_price":selectedItem.Price,
    })
}

func DeleteItemHandler(c *gin.Context) {
    var item model.MenuItem

    // Faz o bind do JSON para o struct item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
        return
    }

    // Chama o serviço para deletar o item
    err := service.DeleteItem(item)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o item: " + err.Error()})
        return
    }

    // Responde com sucesso
    c.JSON(http.StatusOK, gin.H{
        "message": "Item deletado com sucesso!",
    })
}

func UpdateItemHandler(c *gin.Context) {
    var item model.MenuItem

    // Faz o bind do JSON para o struct item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
        return
    }

    // Chama o serviço para deletar o item
    err := service.UpdateItem(item)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o item: " + err.Error()})
        return
    }

    // Responde com sucesso
    c.JSON(http.StatusOK, gin.H{
        "message": "Item atualizado com sucesso!",
    })
}