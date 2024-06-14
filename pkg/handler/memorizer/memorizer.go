package memorizer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/handler/mapper"
)

var destination string = "../../templates/"

func MemorizeTemplate(alias, path string) {
	fmt.Println("path: ", path)
	mapper.SetNewAlias(alias, "caminhos")

	dirPath, errX := os.Getwd()
	if errX != nil {
		log.Fatalln("Error: ", errX)
	}

	completedPath := filepath.Join(dirPath, path)
	completedPath2 := filepath.Join(dirPath, "templates", path)

	fmt.Println("completePath: ", completedPath)

	copyFolder(completedPath, completedPath2)
}
