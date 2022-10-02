package users

import (
	"errors"
	"testing"

	"goLang/pkg/user"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserRequest_IsValid(t *testing.T) {
	type testCase struct {
		request  CreateUserRequest
		expected error
	}

	testCases := map[string]testCase{
		"valid request": {
			request: CreateUserRequest{
				Name:     "test name",
				Email:    "test@email",
				Password: "test password",
			},
			expected: nil,
		},
		"name is empty": {
			request: CreateUserRequest{
				Name:     "",
				Email:    "test@email",
				Password: "test password",
			},
			expected: errors.New("name is empty"),
		},
		"email is empty": {
			request: CreateUserRequest{
				Name:     "test name",
				Email:    "",
				Password: "test password",
			},
			expected: errors.New("email is empty"),
		},
		"email is invalid": {
			request: CreateUserRequest{
				Name:     "test name",
				Email:    "test email",
				Password: "test password",
			},
			expected: errors.New("email is not valid"),
		},
		"password is empty": {
			request: CreateUserRequest{
				Name:     "test name",
				Email:    "test@email",
				Password: "",
			},
			expected: errors.New("password is empty"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.request.IsValid())
		})
	}
}

func TestService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := &CreateUserRequest{
		Name:     "Vasya",
		Email:    "vasya@gmail.com",
		Password: "12345678",
	}

	response := &CreateUserResponse{}

	userRow := &user.User{
		Name:         "Vasya",
		Email:        "vasya@gmail.com",
		PasswordHash: "12345678",
	}

	storageMock := NewMockStorage(ctrl)
	storageMock.EXPECT().Create(userRow).Return(nil)

	service := NewService(storageMock)

	assert.NoError(t, service.CreateUser(nil, request, response))
	assert.Equal(t, &CreateUserResponse{
		User: user.User{
			Name:         "Vasya",
			Email:        "vasya@gmail.com",
			PasswordHash: "12345678",
		},
	}, response)
}
