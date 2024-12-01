package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/nicolasdri92/ms-lib-framework/internal/constants"
)

func LoadEnv() error {
	err := godotenv.Load(".env")

	if err != nil {
		return errors.New(constants.ErrorDotenvNotFound)
	}

	return nil
}

func GetVariable(name string) string {
	return os.Getenv(name)
}
