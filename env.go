package framework

import (
	"errors"
	"os"
	"strconv"

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

func GetVariableBool(name string, defaultValue bool) bool {
	variable := GetVariable(name)

	if value, err := strconv.ParseBool(variable); err == nil {
		return value
	}

	return defaultValue
}

func GetVariableInt(name string, defaultValue int) int {
	variable := GetVariable(name)

	if value, err := strconv.Atoi(variable); err == nil {
		return value
	}

	return defaultValue
}

func GetVariableFloat(name string, defaultValue float64) float64 {
	variable := GetVariable(name)

	if value, err := strconv.ParseFloat(variable, 64); err == nil {
		return value
	}

	return defaultValue
}
