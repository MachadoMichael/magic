package mapper

import (
	"encoding/json"
	"log"
	"os"
)

func UpdateInfoMap(data InfoMap, alias string) error {
	magicMap[alias] = data
	updatedJSON, err := json.MarshalIndent(magicMap, "", " ")
	if err != nil {
		return nil
	}

	err = os.WriteFile(mappingPath, updatedJSON, 0644)
	if err != nil {
		log.Fatal("Cannot write in file")
		return err
	}

	return nil
}
