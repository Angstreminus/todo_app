package controller

import (
	"net/http"
	"strconv"
	"todo_app/model"
	"todo_app/service"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoItemServ *service.TodoItemService
}

func NewTodoController(serv *service.TodoItemService) *TodoController {
	return &TodoController{
		todoItemServ: serv,
	}
}

func (tc TodoController) CreateItem(ctx *gin.Context) {
	var item model.TodoItem

	err := ctx.BindJSON(item)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	resp, respErr := tc.todoItemServ.CreateItem(&item)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (tc TodoController) UpdateItem(ctx *gin.Context) {
	var item model.TodoItem
	err := ctx.BindJSON(&item)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	respErr := tc.todoItemServ.UpdateTodo(&item)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (tc TodoController) FinishItem(ctx *gin.Context) {
	item_id := ctx.Param("id")

	itemId, err := strconv.Atoi(item_id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	respErr := tc.todoItemServ.FinishItem(itemId)

	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
	}

	ctx.Status(http.StatusNoContent)
}

func (tc TodoController) GetItems(ctx *gin.Context) {
	user_id := ctx.Param("id")

	userId, err := strconv.Atoi(user_id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	items, respErr := tc.todoItemServ.GetItems(userId)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
	}

	ctx.JSON(http.StatusOK, items)
}

func (tc TodoController) GetItem(ctx *gin.Context) {
	item_id := ctx.Param("id")

	itemId, err := strconv.Atoi(item_id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	item, respErr := tc.todoItemServ.GetItem(itemId)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
	}

	ctx.JSON(http.StatusOK, item)
}
