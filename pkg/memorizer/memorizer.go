package memorizer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

var destination string = "../../templates/"

func MemorizeTemplate(alias, path string) {
	fmt.Println("path: ", path)
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

	fmt.Println("completePath: ", completedPath)

	archiver.CopyFolder(completedPath, completedPath2)
}
