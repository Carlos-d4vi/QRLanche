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

	r.GET("/listitems", controller.GetAllMenuItemsHandler)
	
	r.GET("/findItem", controller.SelectItemByIdHandler)
	
	r.POST("/newitem", controller.CreateMenuItemHandler)
	
	r.DELETE("/deleteItem", controller.DeleteItemHandler)
	
	r.PUT("/updateItem", controller.UpdateItemHandler)

	// Iniciando o servidor na porta 8080 usando o Gin
	fmt.Println("Iniciando servidor na porta 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
