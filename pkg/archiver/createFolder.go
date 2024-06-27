package archiver

import (
	"fmt"
	"os"
)

func CreateFolder(directoryPath string) error {
	_, err := os.Stat(directoryPath)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
		return nil
	}

	return err
}
