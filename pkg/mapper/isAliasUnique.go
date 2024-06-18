package mapper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func isAliasUnique(alias string) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error to read .env file")
	}

	mapping := os.Getenv("MAPPING_PATH")
	if mapping == "" {
		log.Fatalln("Error to read .env file")
	}

	file, err := os.Open(mapping)
	if err != nil {
		return false, err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return false, err
	}

	var jsonData map[string]InfoMap

	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		return false, err
	}

	_, exist := jsonData[alias]
	fmt.Println(exist)
	return !exist, nil
}
