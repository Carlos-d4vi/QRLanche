package service

import (
	"QrLanche/backend/config"
	"QrLanche/backend/pkg/model"
	"fmt"
)

func CreateMenuItem(item model.MenuItem) (int, error) {
    sqlStatement := `
        INSERT INTO menu_items (name, price)
        VALUES ($1, $2)
        RETURNING id`
    var id int

    // Verificar se DB está inicializado
    if config.DB == nil {
        return 0, fmt.Errorf("banco de dados não inicializado")
    }

    err := config.DB.QueryRow(sqlStatement, item.Name, item.Price).Scan(&id)
    if err != nil {
        return 0, err
    }
    fmt.Printf("Novo pedido com ID: %d criado.\n", id)
    return id, nil
}

func GetAllMenuItems() ([]model.MenuItem, error) {
	sqlStatement := `SELECT id, name, price FROM menu_items`

    rows, err := config.DB.Query(sqlStatement)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar itens do menu: %v", err)
    }
    defer rows.Close()

    var menuItems []model.MenuItem

    for rows.Next() {
        var item model.MenuItem

        // Scaneando os resultados para o struct MenuItem
        err = rows.Scan(&item.ID, &item.Name, &item.Price)
        if err != nil {
            return nil, fmt.Errorf("erro ao ler resultados: %v", err)
        }

        // Adicionando o item ao slice
        menuItems = append(menuItems, item)
    }
    
    // Verificando erros de iteração
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("erro na iteração dos resultados: %v", err)
    }

    return menuItems, nil
}

func SelectItemById(item model.MenuItem) (model.MenuItem, error) {
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

func DeleteItem(item model.MenuItem) error {
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

func UpdateItem(item model.MenuItem) error{
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
