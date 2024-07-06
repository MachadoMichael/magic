package memorizer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func SaveTemplate(alias, path, parameter string) {
	err := mapper.SaveInfoMap(alias, path, parameter)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	dir, errDir := os.Getwd()
	if errDir != nil {
		log.Fatalln("Error: ", errDir)
		return
	}

	exePath, err := os.Executable()
	if err != nil {
		println(err.Error())
	}
	rootPath := filepath.Dir(exePath)

	completedPath := filepath.Join(dir, path)
	dst := filepath.Join(rootPath, "repository", alias, path)

	if parameter == "--file" {
		newFolder := filepath.Join(dir, "repository", alias)
		archiver.CreateFolder(newFolder)
		archiver.CopyFile(completedPath, dst)
	} else {
		err := archiver.CopyFolder(completedPath, dst)
		if err != nil {
			println(err.Error())
		}
	}
}
