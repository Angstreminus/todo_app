package repository

import (
	"database/sql"
	"net/http"
	"todo_app/model"
)

type UserRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserRepository(dbHandler *sql.DB) *UserRepository {
	return &UserRepository{
		dbHandler: dbHandler,
	}
}

func (ur UserRepository) CreateUser(user *model.User) (*model.User, *model.ResponseError) {
	query := `
	INSERT INTO users(first_name, last_name, age, country)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	rows, err := ur.dbHandler.Query(query, user.FirstName, user.LastName, user.Age, user.Country)

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	var userId int

	for rows.Next() {
		err = rows.Scan(userId)
		if err != nil {
			return nil, &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &model.User{
		ID:        userId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Country:   user.Country,
	}, nil
}

func (ur UserRepository) UpdateUser(user *model.User) *model.ResponseError {
	query := `
	UPDATE users
	SET
	first_name = $1,
	last_name = $2,
	age = $3,
	country = $4,
	WHERE id = $5`

	res, err := ur.dbHandler.Exec(query, user.FirstName, user.LastName, user.Age, user.Country, user.ID)

	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAff, err := res.RowsAffected()

	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAff == 0 {
		return &model.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

func (ur UserRepository) DeleteUser(userId int) *model.ResponseError {
	query := `DELETE FROM users
	WHERE $id = $1;`

	res, err := ur.dbHandler.Exec(query, userId)
	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAff, err := res.RowsAffected()

	if err != nil {
		return &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAff == 0 {
		return &model.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

func (ur UserRepository) GetUserByID(userID int) (*model.User, *model.ResponseError) {
	query := `
	SELECT *
	FROM users
	WHERE id = $1`

	var (
		id, age                        int
		first_name, last_name, country string
	)
	rows, err := ur.dbHandler.Query(query, userID)

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(id, first_name, last_name, age, country)
		if err != nil {
			return nil, &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &model.User{
		ID:        userID,
		FirstName: first_name,
		LastName:  last_name,
		Age:       age,
		Country:   country,
	}, nil
}

type Database interface {
	GetUserByID()
	CreateUser()
	UpdateUser()
	DeleteUser()
}
