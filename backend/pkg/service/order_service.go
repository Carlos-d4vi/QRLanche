package service

import (
	"QrLanche/backend/config"
	"QrLanche/backend/pkg/model"
	"fmt"
	"github.com/lib/pq"
)

// Função para criar uma nova ordem
func CreateOrder(order model.Order) (int, error) {
	// Monta a query SQL para inserir a ordem
	sqlStatement := `
		INSERT INTO orders (customer_id, itens, total, table_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	
	var id int

	// Verifica se o banco de dados está inicializado
	if config.DB == nil {
		return 0, fmt.Errorf("banco de dados não inicializado")
	}

	item, err := GetNamesAndPricesByIds(order.Itens)

	total := GetTotalOfPricesByMenuItem(item)

	// Executa a query para criar o pedido
	err = config.DB.QueryRow(sqlStatement, order.CustomerID, pq.Array(order.Itens), total, order.TableID).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Novo pedido com ID: %d criado.\n", id)
	return id, nil
}

// Função para buscar todas as ordens
func GetAllOrders() ([]model.Order, error) {
    sqlStatement := `SELECT id, customer_id, itens, total, table_id FROM orders`

    rows, err := config.DB.Query(sqlStatement)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar ordens: %v", err)
    }
    defer rows.Close()

    var orders []model.Order

    for rows.Next() {
        var order model.Order
        var itens pq.Int64Array // Alteração: usando pq.Int64Array para lidar com o array de IDs
        
        err = rows.Scan(&order.ID, &order.CustomerID, &itens, &order.Total, &order.TableID)
        if err != nil {
            return nil, fmt.Errorf("erro ao ler resultados: %v", err)
        }

        // Converta o pq.Int64Array para []int se necessário
        order.Itens = make([]int, len(itens))
        for i, item := range itens {
            order.Itens[i] = int(item)
        }

        orders = append(orders, order)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("erro na iteração dos resultados: %v", err)
    }

    return orders, nil
}

// Função para buscar uma ordem por ID
func SelectOrderById(id int64) (model.Order, error) {
	sqlStatement := `SELECT id, customer_id, itens, total, table_id FROM orders WHERE id = $1`

	var order model.Order
	var itens pq.Int64Array // Alteração: usando pq.Int64Array para lidar com o array de IDs
	row := config.DB.QueryRow(sqlStatement, id)

	err := row.Scan(&order.ID, &order.CustomerID, &itens, &order.Total, &order.TableID)
	if err != nil {
		return model.Order{}, fmt.Errorf("erro ao encontrar a ordem: %v", err)
	}

	order.Itens = make([]int, len(itens))
	for i, item := range itens {
		order.Itens[i] = int(item)
	}

	return order, nil
}

// Função para deletar uma ordem
func DeleteOrder(id int64) error {
	sqlStatement := `DELETE FROM orders WHERE id = $1`

	_, err := config.DB.Exec(sqlStatement, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar a ordem com ID %d: %v", id, err)
	}

	fmt.Printf("Ordem deletada: ID=%d\n", id)
	return nil
}

// Função para atualizar uma ordem
func UpdateOrder(order model.Order) error {
	sqlStatement := `
	UPDATE orders
	SET customer_id = $1, itens = $2, total = $3,  table_id = $4
	WHERE id = $5`

	item, err := GetNamesAndPricesByIds(order.Itens)
	if err != nil {
		return err
	}

	total := GetTotalOfPricesByMenuItem(item)

	_, err = config.DB.Exec(sqlStatement, order.CustomerID, pq.Array(order.Itens), total, order.TableID, order.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar a ordem com ID %d: %v", order.ID, err)
	}

	fmt.Printf("Ordem atualizada: ID=%d\n", order.ID)
	return nil
}

func GetTotalOfPricesByMenuItem(menuItem []model.MenuItem) (float64){
	var total float64
	for _, price := range menuItem {
		total += price.Price
	}
	return total
}

// Função para buscar os nomes e preços dos itens do menu por seus IDs
func GetNamesAndPricesByIds(ids []int) ([]model.MenuItem, error) {
    sqlStatement := `SELECT id, name, price FROM menu_items WHERE id = ANY($1);`

    rows, err := config.DB.Query(sqlStatement, pq.Array(ids))
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar nomes e preços: %v", err)
    }
    defer rows.Close()

    var menuItems []model.MenuItem
    for rows.Next() {
        var item model.MenuItem
        if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
            return nil, fmt.Errorf("erro ao ler os nomes e preços: %v", err)
        }
        menuItems = append(menuItems, item)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("erro na iteração dos resultados: %v", err)
    }

    return menuItems, nil
}