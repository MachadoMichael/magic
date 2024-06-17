package memorizer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/handler/archiver"
	"github.com/MachadoMichael/magic/pkg/handler/mapper"
)

var destination string = "../../templates/"

func MemorizeTemplate(alias, path string) {
	fmt.Println("path: ", path)
	mapper.SetNewAlias(alias, path)

	dir, errX := os.Getwd()
	if errX != nil {
		log.Fatalln("Error: ", errX)
	}

	completedPath := filepath.Join(dir, path)
	completedPath2 := filepath.Join(dir, "templates", path)

	fmt.Println("completePath: ", completedPath)

	archiver.CopyFolder(completedPath, completedPath2)
}
