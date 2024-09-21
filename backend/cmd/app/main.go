package main

import (
	"QrLanche/backend/config"
	"QrLanche/backend/pkg/controller"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Abrindo a conexão com o banco de dados
	err := config.OpenConn()
	if err != nil {
		log.Printf("Erro ao conectar no banco de dados PostgreSQL: %v", err)
		return
	}
	defer config.CloseConn()
	fmt.Println("Conexão estabelecida com sucesso!")

	// Criando uma nova instância do Gin
	r := gin.Default()
	
	// Menu_items
	r.POST("/newitem", controller.CreateMenuItemHandler)
	
	r.GET("/listitems", controller.GetAllMenuItemsHandler)
	
	r.GET("/finditem", controller.SelectItemByIdHandler)
	
	r.DELETE("/deleteitem", controller.DeleteItemHandler)
	
	r.PUT("/updateitem", controller.UpdateItemHandler)
	
	// Order
	r.POST("/neworder", controller.CreateOrderHandler)
	
	r.GET("/getallorder", controller.GetAllOrdersHandler)
	
	r.GET("/getorderbyid", controller.SelectOrderByIdHandler)
	
	r.PUT("/updateorder", controller.UpdateOrderHandler)

	r.DELETE("/deleteorder", controller.DeleteOrderHandler)
	
	// Customer
	r.POST("/newcustomer", controller.CreateCustomerHandler)
	
	r.GET("/getallcustomers", controller.GetAllCustomersHandler)
	
	r.GET("/getcustomer", controller.GetCustomerByIDHandler)
	
	r.PUT("/updatecustomer", controller.UpdateCustomerHandler)

	r.DELETE("/deletecustomer", controller.DeleteCustomerHandle)

	// RestaurantTable
	r.POST("/newtrestaurantTable", controller.CreateRestaurantTableHandler)
	
	r.GET("/getalltrestaurantTables", controller.GetAllRestaurantTablesHandler)
	
	r.GET("/gettrestaurantTable", controller.GetRestaurantTableByIDHandler)
	
	r.PUT("/updatetrestaurantTable", controller.UpdateRestaurantTableHandler)

	r.DELETE("/deletetrestaurantTable", controller.DeleteRestaurantTableHandler)

	// Iniciando o servidor na porta 8080 usando o Gin
	fmt.Println("Iniciando servidor na porta 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
