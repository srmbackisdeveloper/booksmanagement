package common

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	// check if prod
	prod := os.Getenv("PROD")
	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
