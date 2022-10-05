package users

import (
	"goLang/pkg/user"
	"log"
	"net/http"
)

type FindUserByNameAndEmailRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FindUserByNameAndEmailResponse struct {
	Users []user.User
}

func (h *Service) FindUsersByNameAndEmail(_ *http.Request, req *FindUserByNameAndEmailRequest, res *FindUserByNameAndEmailResponse) error {
	name := req.Name
	email := req.Email
	users, err := h.storage.FindUsersNameEmail(name, email)
	if err != nil {
		log.Println("Can't find user by name and email")
		return err
	}
	res.Users = users
	return nil
}
