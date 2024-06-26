package archiver

import (
	"os"
)

func CopyFile(src, dst string) error {
	println("Copy file: ", src, dst)
	data, err := os.ReadFile(src)
	if err != nil {
		println("Error reading file:", err)
		return err
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		println("Error writing file:", err)
		return err
	}
	return nil
}
