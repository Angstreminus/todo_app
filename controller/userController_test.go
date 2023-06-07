package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_app/model"
	"todo_app/repository"
	"todo_app/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func initTestRouter(dbHandler *sql.DB) *gin.Engine {
	userRepository := repository.NewUserRepository(dbHandler)
	userService := service.NewUserService(userRepository, nil)
	userController := NewUserController(userService)
	router := gin.Default()
	router.GET("/users/:id", userController.GetUser)
	return router
}

func TestGetUserResponse(t *testing.T) {
	dbHandler, mock, _ := sqlmock.New()
	defer dbHandler.Close()
	coulums := []string{"id", "first_name", "last_name", "age", "country"}
	mock.ExpectQuery("SELECT * FROM users WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows(coulums).AddRow(1, "Adiana", "Folkner", 33, "Columbia"))
	router := initTestRouter(dbHandler)
	request, _ := http.NewRequest("GET", "/users/1", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	var user *model.User
	json.Unmarshal(recorder.Body.Bytes(), &user)
	assert.NotEmpty(t, user)
	assert.Equal(t, 33, user.Age)
}
