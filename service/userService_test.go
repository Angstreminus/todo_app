package service

import (
	"net/http"
	"testing"
	"todo_app/model"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserInvalidFirstName(t *testing.T) {
	user := &model.User{
		LastName: "Georg",
		Age:      22,
		Country:  "Canada",
	}

	respErr := validateUser(user)
	assert.NotEmpty(t, respErr)
	assert.Equal(t, "Invalid first name value", respErr.Message)
	assert.Equal(t, http.StatusBadRequest, respErr.Status)
}
func TestValidateUser(t *testing.T) {
	tests := []struct {
		name string
		user *model.User
		want *model.ResponseError
	}{
		{
			name: "Invalid_First_Name",
			user: &model.User{
				LastName: "Goffman",
				Age:      30,
				Country:  "Peru",
			},
			want: &model.ResponseError{
				Message: "Invalid first name value",
				Status:  http.StatusBadRequest,
			},
		},

		{
			name: "Invalid_Last_Name",
			user: &model.User{
				FirstName: "Tom",
				Age:       33,
				Country:   "Singapure",
			},
			want: &model.ResponseError{
				Message: "Invalid last name value",
				Status:  http.StatusBadRequest,
			},
		},

		{
			name: "Invalid_Country",
			user: &model.User{
				LastName:  "Seroide",
				FirstName: "Ann",
				Age:       30,
			},
			want: &model.ResponseError{
				Message: "Invalid country",
				Status:  http.StatusBadRequest,
			},
		},

		{
			name: "Invalid_Age",
			user: &model.User{
				LastName:  "Blackhand",
				FirstName: "Morgan",
				Age:       -3,
				Country:   "Peru",
			},
			want: &model.ResponseError{
				Message: "Invalid age",
				Status:  http.StatusBadRequest,
			},
		},

		{
			name: "Valid_User",
			user: &model.User{
				LastName:  "Blackhand",
				FirstName: "Morgan",
				Age:       55,
				Country:   "USA",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			respErr := validateUser(test.user)
			assert.Equal(t, test.want, respErr)
		})
	}
}
