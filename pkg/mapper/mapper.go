package mapper

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/MachadoMichael/magic/infra"
)

var magicMap map[string]InfoMap
var mappingPath string

func Init() {
	mappingPath = infra.Config.MappingPath

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
