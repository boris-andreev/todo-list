package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv" // remove after docker is used
)

type Config struct {
	DbConnectionString string
}

var config = &Config{}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	config.DbConnectionString = os.Getenv("DbConnectionString")
}

func GetConfig() *Config {
	return config
}
