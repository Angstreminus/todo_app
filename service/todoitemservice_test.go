package service

import (
	"net/http"
	"testing"
	"todo_app/model"

	"github.com/stretchr/testify/assert"
)

func TestValidateItem(t *testing.T) {
	tests := []struct {
		name string
		item *model.TodoItem
		want *model.ResponseError
	}{
		{
			name: "Invalid_Description",
			item: &model.TodoItem{
				UserID: 3,
				Status: "done",
			},
			want: &model.ResponseError{
				Message: "Invalid description",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Invalid_User_Id",
			item: &model.TodoItem{
				Description: "smth",
				Status:      "done",
			},
			want: &model.ResponseError{
				Message: "Invalid foreign key user Id",
				Status:  http.StatusBadRequest,
			},
		},

		{
			name: "Valid_Item",
			item: &model.TodoItem{
				UserID:      12,
				Description: "smth",
				Status:      "done",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			respErr := validateItem(test.item)
			assert.Equal(t, test.want, respErr)
		})
	}
}
