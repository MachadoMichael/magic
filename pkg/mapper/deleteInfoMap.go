package mapper

import (
	"encoding/json"
	"log"
	"os"
)

func DeleteInfoMap(alias string) error {
	delete(magicMap, alias)
	updatedJSON, err := json.MarshalIndent(magicMap, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(mappingPath, updatedJSON, 0644)
	if err != nil {
		log.Fatal("Cannot write in file")
		return err
	}

	return nil
}
