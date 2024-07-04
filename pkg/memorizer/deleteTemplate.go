package memorizer

import (
	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func DeleteTemplate(alias string) error {
	template, err := mapper.GetInfoMap(alias)
	err = archiver.DeleteFolder(template.Path)
	if err != nil {
		return err
	}

	err = mapper.DeleteInfoMap(alias)
	if err != nil {
		return err
	}

	return nil
}
