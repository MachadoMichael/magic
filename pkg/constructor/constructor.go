package constructor

import (
	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func Build(alias, dst string) error {
	resource, err := mapper.GetTemplate(alias)
	if err != nil {
		return err
	}

	resource.Builds++
	archiver.CopyFolder(resource.Path, dst)

	return nil
}
