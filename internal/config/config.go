package config

import (
	"fmt"
	"os"

	"goLang/pkg/repository"

	"github.com/pkg/errors"
)

var ErrEnvironmentNotDeclared = errors.New("environment not declared")

var localDbConfig = repository.Config{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "1234",
	DBName:   "api_go",
	SSLMode:  "disable",
}

func GetAppEnvironment() string {
	if env := os.Getenv("APP_ENV"); env != "" {
		return env
	}
	return "local"
}

func GetDbConfig(env string) (repository.Config, error) {
	if env == "local" {
		return localDbConfig, nil
	}

	return repository.Config{}, errors.Wrap(
		ErrEnvironmentNotDeclared,
		fmt.Sprintf("unknown env %s", env),
	)
}
