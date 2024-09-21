package service

import (
	"QrLanche/backend/config"
	"QrLanche/backend/pkg/model"
	"fmt"
)

// Função para criar um novo cliente
func CreateCustomer(customer model.Customer) (int64, error) {
	sqlStatement := `
		INSERT INTO customers (name, cpf)
		VALUES ($1, $2)
		RETURNING id`

	var id int64

	// Verifica se o banco de dados está inicializado
	if config.DB == nil {
		return 0, fmt.Errorf("banco de dados não inicializado")
	}

	// Executa a query para criar o cliente
	err := config.DB.QueryRow(sqlStatement, customer.Name, customer.Cpf).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("erro ao criar o cliente: %v", err)
	}

	fmt.Printf("Novo cliente com ID: %d criado.\n", id)
	return id, nil
}

// Função para buscar todos os clientes
func GetAllCustomers() ([]model.Customer, error) {
	sqlStatement := `SELECT id, name, cpf FROM customers`

	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar clientes: %v", err)
	}
	defer rows.Close()

	var customers []model.Customer

	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Cpf)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler os resultados: %v", err)
		}
		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro na iteração dos resultados: %v", err)
	}

	return customers, nil
}

// Função para buscar um cliente por ID
func GetCustomerByID(id int64) (model.Customer, error) {
	sqlStatement := `SELECT id, name, cpf FROM customers WHERE id = $1`

	var customer model.Customer
	row := config.DB.QueryRow(sqlStatement, id)

	err := row.Scan(&customer.ID, &customer.Name, &customer.Cpf)
	if err != nil {
		return model.Customer{}, fmt.Errorf("erro ao encontrar o cliente: %v", err)
	}

	return customer, nil
}

// Função para atualizar um cliente
func UpdateCustomer(customer model.Customer) (error) {
	sqlStatement := `
		UPDATE customers
		SET name = $1
		WHERE id = $2`

	_, err := config.DB.Exec(sqlStatement, customer.Name, customer.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar o cliente com ID %d: %v", customer.ID, err)
	}

	fmt.Printf("Cliente atualizado: ID=%d\n", customer.ID)
	return nil
}

// Função para deletar um cliente
func DeleteCustomer(id int64) error {
	sqlStatement := `DELETE FROM customers WHERE id = $1`

	_, err := config.DB.Exec(sqlStatement, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar o cliente com ID %d: %v", id, err)
	}

	fmt.Printf("Cliente deletado: ID=%d\n", id)
	return nil
}
