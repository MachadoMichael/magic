package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	Version        string
	MappingPath    string
	RepositoryPath string
}

var Config *ConfigData

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Config = &ConfigData{
		Version:        os.Getenv("VERSION"),
		MappingPath:    os.Getenv("MAPPING_PATH"),
		RepositoryPath: os.Getenv("REPOSITORY_PATH"),
	}

}
