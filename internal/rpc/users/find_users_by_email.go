package users

import (
	"goLang/pkg/user"
	"log"
	"net/http"
)

type FindUsersByEmailRequest struct {
	Email string `json:"email"`
}

type FindUsersByEmailResponse struct {
	Users []user.User
}

func (h *Service) FindUsersByEmail(_ *http.Request, req *FindUsersByEmailRequest, res *FindUsersByEmailResponse) error {
	email := req.Email
	users, err := h.storage.FindUsersEmail(email)
	if err != nil {
		log.Println("Can not find user by email")
		return err
	}
	res.Users = users
	return nil
}
