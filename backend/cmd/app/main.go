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
	
	r.POST("/newitem", controller.CreateMenuItemHandler)
	
	r.GET("/listitems", controller.GetAllMenuItemsHandler)
	
	r.GET("/finditem", controller.SelectItemByIdHandler)
	
	r.DELETE("/deleteitem", controller.DeleteItemHandler)
	
	r.PUT("/updateitem", controller.UpdateItemHandler)
	
	r.POST("/neworder", controller.CreateOrderHandler)
	
	r.DELETE("/deleteorder", controller.DeleteOrderHandler)
	
	r.GET("/getallorder", controller.GetAllOrdersHandler)
	
	r.GET("/getorderbyid", controller.SelectOrderByIdHandler)
	
	r.PUT("/updateorder", controller.UpdateOrderHandler)

	// Iniciando o servidor na porta 8080 usando o Gin
	fmt.Println("Iniciando servidor na porta 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
