package memorizer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func MemorizeTemplate(alias, path, parameter string) {
	err := mapper.NewAlias(alias, path)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	dir, errDir := os.Getwd()
	if errDir != nil {
		log.Fatalln("Error: ", errDir)
	}

	completedPath := filepath.Join(dir, path)
	dst := filepath.Join(dir, "templates", path)

	if parameter == "--file" {
		archiver.CopyFile(completedPath, dst)
	} else {
		archiver.CopyFolder(completedPath, dst)
	}
}
