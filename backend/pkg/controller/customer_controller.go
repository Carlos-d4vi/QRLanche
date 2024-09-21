package controller

import (
	"QrLanche/backend/pkg/model"
	"QrLanche/backend/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomerHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := service.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cliente criado com sucesso!",
		"id":      id,
	})
}

func GetAllCustomersHandler(c *gin.Context){

	customers, err:= service.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Clientes listados com sucesso!",
	})

	for i := 0; i < len(customers); i++ {
		c.JSON(http.StatusOK, gin.H{
			"Customer_id": customers[i].ID,
			"Customer_name": customers[i].Name,
			"Customer_cpf": customers[i].Cpf,
		})
	}
}

func GetCustomerByIDHandler(c *gin.Context){
	var customer model.Customer

	err:= c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	customer, err = service.GetCustomerByID(customer.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error ao encontrar o customer": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Clientes encontrado com sucesso!",
		"customer_id": customer.ID,
		"customer_name": customer.Name,
		"customer_cpf": customer.Cpf,
	})
}

func UpdateCustomerHandler(c *gin.Context){
	var customer model.Customer

	err:= c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	err = service.UpdateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error ao encontrar o customer": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Clientes atualizado com sucesso!",
	})
}

func DeleteCustomerHandle(c *gin.Context){
	var customer model.Customer

	err:= c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer, err = service.GetCustomerByID(customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error to find": "ocorreu um erro ao encontrar o customer",
			"error": err.Error()})
		return 
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Cliente deletado com sucesso!",
		"customer_id": customer.ID,
		"customer_name": customer.Name,
		"customer_cpf": customer.Cpf,
	})

	service.DeleteCustomer(customer.ID)

	
}
