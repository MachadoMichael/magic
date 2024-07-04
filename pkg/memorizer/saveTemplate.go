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

	completedPath := filepath.Join(dir, path)
	dst := filepath.Join(dir, "repository", alias, path)

	if parameter == "--file" {
		newFolder := filepath.Join(dir, "repository", alias)
		archiver.CreateFolder(newFolder)
		archiver.CopyFile(completedPath, dst)
	} else {
		archiver.CopyFolder(completedPath, dst)
	}
}
