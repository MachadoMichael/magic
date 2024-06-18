package mapper

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

var magicMap map[string]InfoMap
var fullPath string

func InitMapping() {
	filename := "pkg/mapper/internal/mapping.json"
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fullPath = filepath.Join(path, filename)

	file, err := os.Open(fullPath)
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
