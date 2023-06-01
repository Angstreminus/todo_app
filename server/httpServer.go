package server

import (
	"database/sql"
	"log"
	"todo_app/config"
	"todo_app/controller"
	"todo_app/repository"
	"todo_app/service"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	config             *config.Config
	router             *gin.Engine
	userController     *controller.UserController
	todoItemController *controller.TodoController
}

func InitHttpServer(cfg *config.Config, dbhandler *sql.DB) *HttpServer {
	userRepository := repository.NewUserRepository(dbhandler)
	itemRepository := repository.NewTodoItemRepository(dbhandler)
	userService := service.NewUserService(userRepository, itemRepository)
	itemService := service.NewTodoItemService(userRepository, itemRepository)
	userController := controller.NewUserController(userService)
	itemController := controller.NewTodoController(itemService)

	router := gin.Default()
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUser)
	router.PUT("/users/:id", userController.GetUser)
	router.DELETE("/users/:id", userController.DeleteUser)
	router.POST("/users/items", itemController.CreateItem)
	//TODO finish routes

	return &HttpServer{
		config:             cfg,
		router:             router,
		userController:     userController,
		todoItemController: itemController,
	}
}

func (hS HttpServer) Start() {
	err := hS.router.Run(":8080")
	if err != nil {
		log.Fatalf("Error while start up: %v", err)
	}
}
