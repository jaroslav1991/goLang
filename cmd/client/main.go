package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/json"
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

type FindUsersByEmailRequest struct {
	Email string `json:"email"`
}

type FindUsersByEmailResponse struct {
	Users []user.User
}

type FindUsersByNameAndEmailRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FindUsersByNameAndEmailResponse struct {
	Users []user.User
}

func makeRequest(url, method string, args, result any) error {

	message, err := json.EncodeClientRequest(method, args)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.DecodeClientResponse(resp.Body, result)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	url := "http://localhost:8080/"
	args := GetUserByIdRequest{Id: 22}
	var result GetUserByIdResponse

	if err := makeRequest(url, "users.GetUserById", args, &result); err != nil {
		log.Fatal("Can't make a request")
		return
	}

	fmt.Println(result)

	args2 := FindUsersByEmailRequest{Email: "jo%"}
	var result2 FindUsersByEmailResponse

	if err := makeRequest(url, "users.FindUsersByEmail", args2, &result2); err != nil {
		log.Fatal("Can't make a request")
		return
	}
	fmt.Println(result2)

	args3 := FindUsersByNameAndEmailRequest{Name: "va%", Email: "%test%"}
	var result3 FindUsersByNameAndEmailResponse

	if err := makeRequest(url, "users.FindUsersByNameAndEmail", args3, &result3); err != nil {
		log.Fatal("Can't make a request")
		return
	}
	fmt.Println(result3)

}
