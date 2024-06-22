package memorizer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
	"github.com/joho/godotenv"
)

var repositoryFolder string

func MemorizeTemplate(alias, path, parameter string) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalln("Error to read .env file")
	}

	repositoryFolder = os.Getenv("REPOSITORY_PATH")
	if repositoryFolder == "" {
		log.Fatalln("Error to read .env file")
	}

	repoErr := archiver.CreateFolder(repositoryFolder)
	if repoErr != nil {
		log.Fatalln(repoErr)
	}

	err := mapper.NewAlias(alias, path)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	dir, errDir := os.Getwd()
	if errDir != nil {
		log.Fatalln("Error: ", errDir)
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
