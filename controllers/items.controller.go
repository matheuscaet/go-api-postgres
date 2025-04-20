package controller

import (
	"go-api-postgres/model"
	"go-api-postgres/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type itemController struct {
	itemUseCase usecase.ItemUsecase
}

func NewItemController(usecase usecase.ItemUsecase) itemController {
	return itemController{
		itemUseCase: usecase,
	}
}

func (p *itemController) GetItems(ctx *gin.Context) {

	items, err := p.itemUseCase.GetItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (p *itemController) CreateItem(ctx *gin.Context) {

	var item model.Item
	err := ctx.BindJSON(&item)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedItem, err := p.itemUseCase.CreateItem(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedItem)
}

func (p *itemController) GetItemById(ctx *gin.Context) {

	id := ctx.Param("itemId")
	if id == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	itemId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id should be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	item, err := p.itemUseCase.GetItemById(itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if item == nil {
		response := model.Response{
			Message: "Item not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, item)
}
