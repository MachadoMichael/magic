package mapper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func NewAlias(alias, path string) error {
	exist, err := isAliasUnique(alias)
	if err != nil {
		fmt.Println("here")
		return err
	}

	println("exist ", exist)
	// if !exist {
	// 	return errors.New("already exist this alias")
	// }

	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	info := InfoMap{
		path,
		formattedTime,
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
