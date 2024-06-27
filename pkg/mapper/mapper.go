package mapper

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var magicMap map[string]InfoMap
var mappingPath string

func InitMapping() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalln("Error to read .env file")
	}

	mappingPath = os.Getenv("MAPPING_PATH")
	if mappingPath == "" {
		log.Fatalln("Error to read .env file")
	}

	file, err := os.Open(mappingPath)
	if err != nil {
		log.Fatal("Error on open file", err)
	}

	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("Error when try read file.")
		return
	}

	err = json.Unmarshal(jsonData, &magicMap)
	if err != nil {
		log.Fatalln("Error when try convert file to map.", err)
		return
	}
}
