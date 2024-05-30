package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var destination string = "../../templates/"

func MemorizeTemplate(templatePath, templateName string) {
	baseFolder := filepath.Base(templatePath)
	fmt.Printf("baseFolder: %v", baseFolder)

	files, err := os.ReadDir(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	newProjectPath := filepath.Join(templateName)
	err = os.Mkdir(newProjectPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		srcFile := filepath.Join(templatePath, file.Name())

		err = copyFile(srcFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func copyFile(src string) error {
	fmt.Printf("%v %v ", src, destination)
	return nil
}
