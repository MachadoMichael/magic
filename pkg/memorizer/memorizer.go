package memorizer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

var destination string = "../../templates/"

func MemorizeTemplate(alias, path string) {
	err := mapper.NewAlias(alias, path)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	dir, errDir := os.Getwd()
	if errDir != nil {
		log.Fatalln("Error: ", errDir)
	}

	completedPath := filepath.Join(dir, path)
	completedPath2 := filepath.Join(dir, "templates", path)

	archiver.CopyFolder(completedPath, completedPath2)
}
