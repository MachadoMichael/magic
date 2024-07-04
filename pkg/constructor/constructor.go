package constructor

import (
	"path/filepath"

	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func Build(alias, dst string) error {
	template, err := mapper.GetInfoMap(alias)
	if err != nil {
		return err
	}

	template.Builds++

	if template.IsFile {
		src := filepath.Join(template.Path, template.Resource)
		dst = filepath.Join(dst, template.Resource)
		err = archiver.CopyFile(src, dst)
		if err != nil {
			return err
		}

		return nil
	}

	err = archiver.CopyFolder(template.Path, dst)
	if err != nil {
		println(err.Error(), "copy")
		return err
	}

	return nil
}
