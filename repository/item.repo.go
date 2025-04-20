package repository

import (
	"database/sql"
	"fmt"
	"go-api-postgres/model"
)

type ItemRepository struct {
	connection *sql.DB
}

func NewItemRepository(connection *sql.DB) ItemRepository {
	return ItemRepository{
		connection: connection,
	}
}

func (pr *ItemRepository) GetItems() ([]model.Item, error) {

	query := "SELECT id, name, description FROM item"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Item{}, err
	}

	var itemList []model.Item
	var itemObj model.Item

	for rows.Next() {
		err = rows.Scan(
			&itemObj.ID,
			&itemObj.Name,
			&itemObj.Description)

		if err != nil {
			fmt.Println(err)
			return []model.Item{}, err
		}

		itemList = append(itemList, itemObj)
	}

	rows.Close()

	return itemList, nil
}

func (pr *ItemRepository) CreateItem(item model.Item) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO item (name, description) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(item.Name, item.Description).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ItemRepository) GetItemById(id_item int) (*model.Item, error) {

	query, err := pr.connection.Prepare("SELECT * FROM item WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Item

	err = query.QueryRow(id_item).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}
