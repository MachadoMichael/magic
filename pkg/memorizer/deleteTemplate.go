package memorizer

import (
	"github.com/MachadoMichael/magic/pkg/archiver"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func DeleteTemplate(alias string) error {
	err := mapper.DeleteAlias(alias)
	if err != nil {
		return err
	}

	err = archiver.DeleteFolder(alias)
	if err != nil {
		return err
	}

	return nil
}
