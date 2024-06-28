package memorizer

import (
	"log"

	"github.com/MachadoMichael/magic/infra"
	"github.com/MachadoMichael/magic/pkg/archiver"
)

var repositoryFolder string

func Init() error {
	repositoryFolder = infra.Config.RepositoryPath
	err := archiver.CreateFolder(repositoryFolder)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
