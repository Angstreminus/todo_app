package service

import (
	"net/http"
	"todo_app/model"
	"todo_app/repository"
)

type TodoItemService struct {
	userRep     *repository.UserRepository
	todoItemRep *repository.TodoItemRepository
}

func NewTodoItemService(urRep *repository.UserRepository, todoItemRep *repository.TodoItemRepository) *TodoItemService {
	return &TodoItemService{
		userRep:     urRep,
		todoItemRep: todoItemRep,
	}
}

func (ti TodoItemService) CreateItem(item *model.TodoItem) (*model.TodoItem, *model.ResponseError) {
	respErr := validateItem(item)
	if respErr != nil {
		return nil, respErr
	}

	user, err := ti.userRep.GetUserByID(item.ID)
	if err != nil {
		return nil, respErr
	}

	if user == nil {
		return nil, &model.ResponseError{
			Message: "User not found!",
			Status:  http.StatusNotFound,
		}
	}
	//TODO controller USER_ID foreign key check
	return ti.todoItemRep.CreateTodoItem(*item)
}

func (ts TodoItemService) UpdateTodo(todo *model.TodoItem) *model.ResponseError {
	respErr := validateItem(todo)
	if respErr != nil {
		return respErr
	}
	return nil
}

func (ts TodoItemService) FinishItem(itemId int) *model.ResponseError {
	respErr := validateItemId(itemId)

	if respErr != nil {
		return respErr
	}

	return ts.todoItemRep.FinishItem(itemId)
}

func (ts TodoItemService) GetItems(userId int) ([]*model.TodoItem, *model.ResponseError) {
	respErr := validateItemId(userId)
	if respErr != nil {
		return nil, respErr
	}

	return ts.todoItemRep.GetItems(userId)
}

func (ts TodoItemService) GetItem(itemId int) (*model.TodoItem, *model.ResponseError) {
	respErr := validateItemId(itemId)
	if respErr != nil {
		return nil, respErr
	}

	return ts.todoItemRep.GetItem(itemId)
}

func validateItem(item *model.TodoItem) *model.ResponseError {
	if item.Description == "" {
		return &model.ResponseError{
			Message: "Invalid description",
			Status:  http.StatusBadRequest,
		}
	}

	if item.UserID == 0 {
		return &model.ResponseError{
			Message: "Invalid foreign key user Id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func validateItemId(itemId int) *model.ResponseError {
	if itemId == 0 {
		return &model.ResponseError{
			Message: "Invalid id",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
