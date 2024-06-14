package mapper

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type InfoMap struct {
	Path     string `json:"path"`
	CreateAt string `json:"createAt"`
}

var magicMap map[string]InfoMap
var fullPath string

func InitMapping() {
	filename := "pkg/handler/mapper/internal/mapping.json"
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

func SetNewAlias(alias, path string) {
	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	info := InfoMap{
		path,
		formattedTime,
	}

	magicMap[alias] = info
	updatedJSON, err := json.MarshalIndent(magicMap, "", " ")
	if err != nil {
		return
	}

	err = os.WriteFile(fullPath, updatedJSON, 0644)
	if err != nil {
		log.Fatal("Cannot write in file")
		return
	}

}
