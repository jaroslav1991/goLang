package main

import (
	"log"

	"goLang"
	"goLang/internal/config"
	"goLang/pkg/handler"
	"goLang/pkg/repository"
	"goLang/pkg/service"
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

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.HewHandler(services)

	srv := new(goLang.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server %s", err.Error())
	}
}
