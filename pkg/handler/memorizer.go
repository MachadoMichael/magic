package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func MemorizeTemplate(templatePath, templateName string) {
	baseFolder := filepath.Base(templatePath)
	fmt.Println(baseFolder)

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
		dstFile := filepath.Join(templatePath, file.Name())

		err = copyFile(srcFile, dstFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func copyFile(src, dst string) error {
	return nil
}
