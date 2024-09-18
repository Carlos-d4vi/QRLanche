package service

import (
	"QrLanche/backend/config"
	"QrLanche/backend/pkg/model"
	"fmt"

	"github.com/lib/pq"
)

func CreateOrder(order model.Order) (int, error) {
    sqlStatement := `
        INSERT INTO orders (customers, itens, total)
        VALUES ($1, $2, $3)
        RETURNING id`
    var id int

    // Verificar se DB está inicializado
    if config.DB == nil {
        return 0, fmt.Errorf("banco de dados não inicializado")
    }

    err := config.DB.QueryRow(sqlStatement, order.Customer, order.Itens).Scan(&id)
    if err != nil {
        return 0, err
    }
    fmt.Printf("Novo pedido com ID: %d criado.\n", id)
    return id, nil
}

func GetAllOrdes() ([]model.MenuItem, error) {
	sqlStatement := `SELECT id, name, price FROM menu_items`

    rows, err := config.DB.Query(sqlStatement)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar itens do menu: %v", err)
    }
    defer rows.Close()

    var Order []model.MenuItem

    for rows.Next() {
        var item model.MenuItem

        // Scaneando os resultados para o struct MenuItem
        err = rows.Scan(&item.ID, &item.Name, &item.Price)
        if err != nil {
            return nil, fmt.Errorf("erro ao ler resultados: %v", err)
        }

        // Adicionando o item ao slice
        Order = append(Order, item)
    }
    
    // Verificando erros de iteração
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("erro na iteração dos resultados: %v", err)
    }

    return Order, nil
}

func SelectOrderById(item model.MenuItem) (model.MenuItem, error) {
    sqlStatement := `SELECT id, name, price FROM menu_items WHERE id = $1;`

    var selectedItem model.MenuItem
    row := config.DB.QueryRow(sqlStatement, item.ID)

    // Escaneia os resultados e armazena no selectedItem
    err := row.Scan(&selectedItem.ID, &selectedItem.Name, &selectedItem.Price)
    if err != nil {
        return model.MenuItem{}, fmt.Errorf("erro ao encontrar o item: %v", err)
    }

    // Retorna o item selecionado
    return selectedItem, nil
}

func DeleteOrder(item model.MenuItem) error {
    sqlStatement := `DELETE FROM menu_items WHERE id = $1;`

    // Executa a query para deletar o item
    _, err := config.DB.Exec(sqlStatement, item.ID)
    if err != nil {
        return fmt.Errorf("erro ao deletar o item com ID %d: %v", item.ID, err)
    }

    // Mensagem de sucesso
    fmt.Printf("Item deletado: ID=%d, Name=%s, Price=%.2f\n", item.ID, item.Name, item.Price)
    return nil
}

func UpdateOrder(item model.MenuItem) error{
    sqlStatement := `
    UPDATE menu_items
    SET name = $1, price = $2
    WHERE id = $3`
    _, err := config.DB.Exec(sqlStatement, item.Name, item.Price, item.ID)
    if err != nil {
        return fmt.Errorf("erro ao atualizar o item com ID %d: %v", item.ID, err)
    }

    fmt.Printf("Item atualizado: ID=%d, Name=%s, Price=%.2f\n", item.ID, item.Name, item.Price)
    return nil
}

func GetPricesByIds(ids []int) ([]float64, error) {
    // Constrói a query SQL
    sqlStatement := `SELECT price FROM menu_items WHERE id = ANY($1);`

    // Executa a query
    rows, err := config.DB.Query(sqlStatement, pq.Array(ids))
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar preços: %v", err)
    }
    defer rows.Close()

    var prices []float64
    for rows.Next() {
        var price float64
        if err := rows.Scan(&price); err != nil {
            return nil, fmt.Errorf("erro ao ler os preços: %v", err)
        }
        prices = append(prices, price)
    }

    return prices, nil
}
