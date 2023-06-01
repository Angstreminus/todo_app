package controller

import (
	"net/http"
	"strconv"
	"todo_app/model"
	"todo_app/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(usrServ *service.UserService) *UserController {
	return &UserController{
		userService: usrServ,
	}
}

func (uc UserController) CreateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.BindJSON(user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	resp, respErr := uc.userService.CreateUser(&user)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (uc UserController) UpdateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	respErr := uc.userService.UpdateUser(&user)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (uc UserController) DeleteUser(ctx *gin.Context) {
	user_id := ctx.Param("id")

	userId, err := strconv.Atoi(user_id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	respErr := uc.userService.DeleteUser(userId)

	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
	}

	ctx.Status(http.StatusNoContent)
}

func (uc UserController) GetUser(ctx *gin.Context) {
	user_id := ctx.Param("id")

	userId, err := strconv.Atoi(user_id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	user, respErr := uc.userService.GetUserByID(userId)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
	}

	ctx.JSON(http.StatusOK, user)
}
