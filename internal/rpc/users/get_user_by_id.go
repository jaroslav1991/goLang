package users

import (
	"errors"
	"goLang/pkg/user"
	"log"
	"net/http"
)

type GetUserByIdRequest struct {
	Id int `json:"id"`
}

type GetUserByIdResponse struct {
	User user.User
}

func (h *Service) GetUserById(_ *http.Request, req *GetUserByIdRequest, res *GetUserByIdResponse) error {
	id := req.Id
	usr, err := h.storage.GetUser(id)
	if err != nil {
		log.Printf("get user by %d id failed with error", id)
		return err
	}
	if usr == nil {
		return errors.New("user not found")
	}

	res.User = *usr
	return nil
}
