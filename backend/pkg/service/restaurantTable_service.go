package service

import (
	"QrLanche/backend/config"
	"QrLanche/backend/pkg/model"
	"fmt"
)

// Função para criar uma nova mesa
func CreateRestaurantTable(table model.RestaurantTable) (int64, error) {
	sqlStatement := `
		INSERT INTO restaurant_tables (number)
		VALUES ($1)
		RETURNING id`

	var id int64

	// Verifica se o banco de dados está inicializado
	if config.DB == nil {
		return 0, fmt.Errorf("banco de dados não inicializado")
	}

	// Executa a query para criar a mesa
	err := config.DB.QueryRow(sqlStatement, table.Number).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("erro ao criar a mesa: %v", err)
	}

	fmt.Printf("Nova mesa com ID: %d criada.\n", id)
	return id, nil
}

// Função para buscar todas as mesas
func GetAllRestaurantTables() ([]model.RestaurantTable, error) {
	sqlStatement := `SELECT id, number, available FROM restaurant_tables`

	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar mesas: %v", err)
	}
	defer rows.Close()

	var tables []model.RestaurantTable

	for rows.Next() {
		var table model.RestaurantTable
		err = rows.Scan(&table.ID, &table.Number, &table.Available)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler os resultados: %v", err)
		}
		tables = append(tables, table)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro na iteração dos resultados: %v", err)
	}

	return tables, nil
}

// Função para buscar uma mesa por ID
func GetRestaurantTableByID(id int64) (model.RestaurantTable, error) {
	sqlStatement := `SELECT id, number, available FROM restaurant_tables WHERE id = $1`

	var table model.RestaurantTable
	row := config.DB.QueryRow(sqlStatement, id)

	err := row.Scan(&table.ID, &table.Number, &table.Available)
	if err != nil {
		return model.RestaurantTable{}, fmt.Errorf("erro ao encontrar a mesa: %v", err)
	}

	return table, nil
}

// Função para atualizar uma mesa
func UpdateRestaurantTable(table model.RestaurantTable) error {
	sqlStatement := `
		UPDATE restaurant_tables
		SET number = $1, available = $2
		WHERE id = $3`

	_, err := config.DB.Exec(sqlStatement, table.Number, table.Available, table.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar a mesa com ID %d: %v", table.ID, err)
	}

	fmt.Printf("Mesa atualizada: ID=%d\n", table.ID)
	return nil
}

// Função para deletar uma mesa
func DeleteRestaurantTable(id int64) error {
	sqlStatement := `DELETE FROM restaurant_tables WHERE id = $1`

	_, err := config.DB.Exec(sqlStatement, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar a mesa com ID %d: %v", id, err)
	}

	fmt.Printf("Mesa deletada: ID=%d\n", id)
	return nil
}