package mapper

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/MachadoMichael/magic/infra"
)

func SaveInfoMap(alias, rsc, parameter string) error {
	exist, err := isAliasUnique(alias)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("already exist this alias")
	}

	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	var isFile bool
	if parameter == "" {
		isFile = false
	} else {
		isFile = true
	}

	path := filepath.Join(infra.Config.RepositoryPath, alias)

	info := InfoMap{
		rsc,
		path,
		formattedTime,
		0,
		isFile,
	}

	magicMap[alias] = info
	updatedJSON, err := json.MarshalIndent(magicMap, "", " ")
	if err != nil {
		return nil
	}

	err = os.WriteFile(mappingPath, updatedJSON, 0644)
	if err != nil {
		log.Fatal("Cannot write in file")
		return nil
	}

	return nil
}
