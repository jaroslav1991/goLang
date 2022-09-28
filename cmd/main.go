package main

import (
	"goLang"
	"goLang/pkg/handler"
	"goLang/pkg/repository"
	"goLang/pkg/service"
	"log"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "1234",
		DBName:   "api_go",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Failed on initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.HewHandler(services)

	srv := new(goLang.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server %s", err.Error())
	}
}
