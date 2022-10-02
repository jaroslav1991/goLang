package main

import (
	"github.com/gorilla/mux"
	"goLang/pkg/user/storage"
	"log"

	"goLang"
	"goLang/internal/config"
	"goLang/pkg/handler"
	"goLang/pkg/repository"
)

func main() {
	env := config.GetAppEnvironment()
	log.Println("current env", env)

	dbConf, err := config.GetDbConfig(env)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.NewPostgresDB(dbConf)
	if err != nil {
		log.Fatalln("Failed on initialize db:", err)
	}

	storage := storage.NewStorage(db)
	handler := handler.NewHandler(storage)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users/create", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/get", handler.GetUsers).Methods("GET")
	router.HandleFunc("/user/get", handler.GetUserById).Methods("GET")

	srv := &goLang.Server{}
	if err := srv.Run("8080", router); err != nil {
		log.Fatalf("error while running http server %s", err.Error())
	}
}
