package archiver

import (
	"fmt"
	"os"
)

func DeleteFolder(directoryPath string) error {
	err := os.RemoveAll(directoryPath)
	if err != nil {
		return fmt.Errorf("error deleting directory: %v", err)
	}
	return nil
}
