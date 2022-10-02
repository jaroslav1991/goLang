//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock_$GOFILE
package users

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"goLang/pkg/user"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

func (r *CreateUserRequest) IsValid() error {
	if r.Name == "" {
		return errors.New("name is empty")
	}

	if r.Email == "" {
		return errors.New("email is empty")
	}

	if strings.Index(r.Email, "@") == -1 {
		return errors.New("email is not valid")
	}

	if r.Password == "" {
		return errors.New("password is empty")
	}

	return nil
}

type CreateUserResponse struct {
	User user.User
}

func (h *Service) CreateUser(_ *http.Request, req *CreateUserRequest, res *CreateUserResponse) error {
	// валидируем запрос
	if err := req.IsValid(); err != nil {
		log.Println("invalid request", err)
		return err
	}

	// готовим пользователя для создания
	userRow := &user.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	// создаем пользователя
	if err := h.storage.Create(userRow); err != nil {
		log.Println("create user failed with error", err)
		return errors.New("create user failed")
	}

	res.User = *userRow
	return nil
}
