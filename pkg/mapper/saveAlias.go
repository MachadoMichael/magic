package mapper

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"
)

func SaveAlias(alias, path string) error {
	exist, err := isAliasUnique(alias)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("already exist this alias")
	}

	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	info := InfoMap{
		path,
		formattedTime,
		0,
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
