package main

import (
	"log"
	"net/http"

	"goLang/internal/config"
	"goLang/internal/rpc/users"
	"goLang/pkg/repository"
	"goLang/pkg/user/storage"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
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

	srv := rpc.NewServer()
	srv.RegisterCodec(json.NewCodec(), "application/json")

	if err := srv.RegisterService(users.NewService(storage), "users"); err != nil {
		log.Fatalln("register createUserHandler error:", err)
	}

	http.Handle("/", srv)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}

	//handler := handler.NewHandler(storage)
	//
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/users/create", handler.CreateUser).Methods("POST")
	//router.HandleFunc("/users/get", handler.GetUsers).Methods("GET")
	//router.HandleFunc("/user/get", handler.GetUserById).Methods("GET")
	//
	//srv := &goLang.Server{}
	//if err := srv.Run("8080", router); err != nil {
	//	log.Fatalf("error while running http server %s", err.Error())
	//}
}
