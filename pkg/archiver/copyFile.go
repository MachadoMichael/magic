package archiver

import (
	"fmt"
	"os"
)

func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}
