package memorizer

import (
	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func DeleteTemplate(alias string) error {
	template, err := mapper.GetTemplate(alias)
	err = archiver.DeleteFolder(template.Path)
	if err != nil {
		return err
	}

	err = mapper.DeleteAlias(alias)
	if err != nil {
		return err
	}

	return nil
}
