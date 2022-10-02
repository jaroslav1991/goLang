package handler

import (
	"encoding/json"
	"goLang/pkg/user"
	"io"
	"log"
	"net/http"
)

type Storage interface {
	Create(user *user.User) error
	GetUser(id int) (*user.User, error)
}

type handler struct {
	storage Storage
}

func NewHandler(storage Storage) *handler {
	return &handler{storage: storage}
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
}

type getUserByIdRequest struct {
	Id int `json:"id"`
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	var request getUserByIdRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("read request body failed with error", err)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("read request body failed with error", err)
		return
	}

	usr, err := h.storage.GetUser(request.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error from getUser on db", err)
		return
	}
	if usr == nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("user not found", err)
		return
	}

	responseBody, _ := json.Marshal(usr)
	if _, err := w.Write(responseBody); err != nil {
		log.Println("write response body failed with error", err)
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRow user.User

	// читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("read request body failed with error", err)
		return
	}

	// парсим запрос из json
	if err := json.Unmarshal(body, &userRow); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("read request body failed with error", err)
		return
	}

	// валидируем запрос
	if userRow.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		if _, err := w.Write([]byte(`{"error": "email is empty"}`)); err != nil {
			log.Println("write response body failed with error", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		log.Println("read request body failed with error", err)
		return
	}

	// создаем пользователя
	if err := h.storage.Create(&userRow); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("create user failed with error", err)
		return
	}

	log.Println("new user id in handler", userRow.UserId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseBody, _ := json.Marshal(userRow)

	if _, err := w.Write(responseBody); err != nil {
		log.Println("write response body failed with error", err)
	}
}

func (h *handler) deleteUser(w http.ResponseWriter, r *http.Request) {
}
