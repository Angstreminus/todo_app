package repository

import (
	"database/sql"
	"net/http"
	"todo_app/model"
)

type TodoItemRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewTodoItemRepository(dbHand *sql.DB) *TodoItemRepository {
	return &TodoItemRepository{
		dbHandler: dbHand,
	}
}

func (tr TodoItemRepository) CreateTodoItem(todo model.TodoItem) (*model.TodoItem, *model.ResponseError) {
	query := `
	INSERT INTO todo(user_id, descr, status)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	todo.Status = "active"
	rows, err := tr.transaction.Query(query, todo.UserID, todo.Description, todo.Status)

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	var todoId int

	for rows.Next() {
		err = rows.Scan(todoId)
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

	return &model.TodoItem{
		ID:          todoId,
		UserID:      todo.UserID,
		Description: todo.Description,
		Status:      todo.Status,
	}, nil
}

func (tr TodoItemRepository) UpdateTodo(todo *model.TodoItem) *model.ResponseError {
	query := `
	UPDATE items
	SET
	user_id = $1,
	descr = $2,
	status = $3,
	WHERE id = $4`

	res, err := tr.dbHandler.Exec(query, todo.UserID, todo.Description, todo.Status)

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
			Message: "TodoItem not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (tr TodoItemRepository) FinishItem(itemId string) *model.ResponseError {
	query := `
	UPDATE users
	SET
	status = $1,
	WHERE id = $2`

	var status = "done"
	res, err := tr.dbHandler.Exec(query, status, itemId)
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
			Message: "TodoItem not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (tr TodoItemRepository) GetItems(userId int) ([]*model.TodoItem, *model.ResponseError) {
	query := `
		SELECT *
		FROM items
	WHERE user_id = $1`

	rows, err := tr.dbHandler.Query(query, userId)
	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	items := make([]*model.TodoItem, 0)
	var (
		id                  int
		description, status string
	)

	for rows.Next() {
		err := rows.Scan(&id, &description, &status)
		if err != nil {
			return nil, &model.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		item := &model.TodoItem{
			ID:          id,
			UserID:      userId,
			Description: description,
			Status:      status,
		}
		items = append(items, item)
	}
	if rows.Err() != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return items, nil
}
