package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var destination string = "../../templates/"

func MemorizeTemplate(alias, path string) {
	fmt.Printf("path: %s", path)

	f1("internal/mapping.json")
	// setNewAlias(alias, "caminhos")

	// err := os.WriteFile(path+alias, []byte("hello, Gophers!"), 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//

	// files, err := os.ReadDir(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// newProjectPath := filepath.Join(alias)
	// err = os.Mkdir(newProjectPath, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, file := range files {
	// 	srcFile := filepath.Join(path, file.Name())
	//
	// 	err = copyFile(srcFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}

func copyFile(src string) error {
	fmt.Printf("%v %v ", src, destination)
	return nil
}

type InfoMap struct {
	Path     string    `json:"path"`
	CreateAt time.Time `json:"createAt"`
}

var magicMap map[string]InfoMap
var fullPath string

func f1(filename string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("--------------")
	fmt.Println(time.Now())
	fmt.Println("--------------")
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

	fmt.Printf("%s", jsonData)

	err = json.Unmarshal(jsonData, &magicMap)
	if err != nil {
		log.Fatalln("Error when try convert file to map.", err)
		return
	}

}

func setNewAlias(alias, path string) {
	info := InfoMap{
		path,
		time.Now(),
	}

	fmt.Println("alias: ", alias)
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
