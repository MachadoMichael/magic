package mapper

import (
	"encoding/json"
	"io"
	"os"
)

func isAliasUnique(alias string) (bool, error) {
	file, err := os.Open(mappingPath)
	if err != nil {
		return false, err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return false, err
	}

	var jsonData map[string]InfoMap

	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		return false, err
	}

	_, exist := jsonData[alias]
	return !exist, nil
}
