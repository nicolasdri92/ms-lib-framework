package framework

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")

	if err != nil {
		return errors.New(ErrorDotenvNotFound)
	}

	return nil
}

func GetVariable(name string) string {
	return os.Getenv(name)
}
