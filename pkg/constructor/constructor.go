package constructor

import (
	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func Build(alias, dst string) error {
	resource, err := mapper.GetTemplate(alias)
	println("resource", resource.Path)
	println("destino", dst)
	if err != nil {
		return err
	}

	resource.Builds++

	if resource.IsFile {
		println("FILE", resource.Path)
		err = archiver.CopyFile(resource.Path, dst)
		if err != nil {
			return err
		}

		println("FILE")
		return nil
	}

	err = archiver.CopyFolder(resource.Path, dst)
	if err != nil {
		println(err, "copy")
		return err
	}

	return nil
}
