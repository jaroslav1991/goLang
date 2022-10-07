package users

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"goLang/pkg/user"
	"testing"
)

func TestService_FindUsersByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	request := &FindUsersByEmailRequest{
		Email: "jo%",
	}

	response := &FindUsersByEmailResponse{}

	user1Row := user.User{
		UserId:       1,
		Name:         "jora",
		Email:        "jora@gmail.com",
		PasswordHash: "1234",
	}
	user2Row := user.User{
		UserId:       2,
		Name:         "vasya",
		Email:        "vasya@gmail.com",
		PasswordHash: "1234",
	}

	email := request.Email

	data := []user.User{user1Row, user2Row}

	storageMock := NewMockStorage(ctrl)
	storageMock.EXPECT().FindUsersEmail(email).Return(data, nil)

	service := NewService(storageMock)
	assert.NoError(t, service.FindUsersByEmail(nil, request, response))
	assert.Equal(t, &FindUsersByEmailResponse{Users: data}, response)

}

func TestService_FindUsersByEmailError(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	request := &FindUsersByEmailRequest{
		Email: "jo%",
	}

	response := &FindUsersByEmailResponse{}

	err := errors.New("trouble")

	storageMock := NewMockStorage(ctrl)
	service := NewService(storageMock)

	storageMock.EXPECT().FindUsersEmail(request.Email).Return(nil, err)
	assert.Error(t, service.FindUsersByEmail(nil, request, response))
}
