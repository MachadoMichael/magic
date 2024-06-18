package mapper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func isAliasUnique(alias, jsonPath string) (bool, error) {
	file, err := os.Open(jsonPath)
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
	fmt.Println(exist)
	return !exist, nil
}
