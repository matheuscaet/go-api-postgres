package usecase

import (
	"go-api-postgres/model"
	"go-api-postgres/repository"
)

type ItemUsecase struct {
	repository repository.ItemRepository
}

func NewItemUseCase(repo repository.ItemRepository) ItemUsecase {
	return ItemUsecase{
		repository: repo,
	}
}

func (pu *ItemUsecase) GetItems() ([]model.Item, error) {
	return pu.repository.GetItems()
}

func (pu *ItemUsecase) CreateItem(item model.Item) (model.Item, error) {

	itemId, err := pu.repository.CreateItem(item)
	if err != nil {
		return model.Item{}, err
	}

	item.ID = itemId

	return item, nil
}

func (pu *ItemUsecase) GetItemById(id_item int) (*model.Item, error) {

	item, err := pu.repository.GetItemById(id_item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
