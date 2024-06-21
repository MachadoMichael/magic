package memorizer

import (
	"fmt"
	"os"
)

func initRepository(directoryPath string) error {
	// Check if the directory exists
	_, err := os.Stat(directoryPath)
	if err == nil {
		// Directory exists, no need to create it
		return nil
	}

	// Directory does not exist, create it
	if os.IsNotExist(err) {
		err = os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
		return nil
	}

	// Some other error occurred
	return err
}
