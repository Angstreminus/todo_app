package service

import (
	"net/http"
	"todo_app/model"
	"todo_app/repository"
)

type UserService struct {
	userRepository     *repository.UserRepository
	todoItemRepository *repository.TodoItemRepository
}

func NewUserService(userRep *repository.UserRepository, todoItemRep *repository.TodoItemRepository) *UserService {
	return &UserService{
		userRepository:     userRep,
		todoItemRepository: todoItemRep,
	}
}

func (us UserService) CreateUser(user *model.User) (*model.User, *model.ResponseError) {
	err := validateUser(user)
	if err != nil {
		return nil, err
	}

	return us.userRepository.CreateUser(user)
}

func (us UserService) UpdateUser(user *model.User) *model.ResponseError {
	respErr := validateUserId(user.ID)
	if respErr != nil {
		return respErr
	}

	respErr = validateUser(user)
	if respErr != nil {
		return respErr
	}

	return us.userRepository.UpdateUser(user)
}

func (us UserService) DeleteUser(userId int) *model.ResponseError {
	respErr := validateUserId(userId)

	if respErr != nil {
		return respErr
	}

	return us.userRepository.DeleteUser(userId)
}

func (us UserService) GetUserByID(userId int) (*model.User, *model.ResponseError) {
	respErr := validateUserId(userId)

	if respErr != nil {
		return nil, respErr
	}

	return us.userRepository.GetUserByID(userId)
}

func validateUserId(userId int) *model.ResponseError {
	if userId == 0 {
		return &model.ResponseError{
			Message: "Invalid user id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func validateUser(user *model.User) *model.ResponseError {
	if user.Age < 0 || user.Age > 200 {
		return &model.ResponseError{
			Message: "Invalid age",
			Status:  http.StatusBadRequest,
		}
	}

	if user.FirstName == "" || user.FirstName == " " || len(user.FirstName) > 1 {
		return &model.ResponseError{
			Message: "Invalid first name value",
			Status:  http.StatusBadRequest,
		}
	}

	if user.LastName == "" || user.LastName == " " || len(user.LastName) > 1 {
		return &model.ResponseError{
			Message: "Invalid last name value",
			Status:  http.StatusBadRequest,
		}
	}

	if user.Country == "" || user.Country == " " {
		return &model.ResponseError{
			Message: "Invalid country",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
