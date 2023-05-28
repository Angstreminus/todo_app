package server

import (
	"todo_app/config"
	"todo_app/controller"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	config             *config.Config
	router             *gin.Engine
	userController     *controller.UserController
	todoItemController *controller.todoItemController
}
