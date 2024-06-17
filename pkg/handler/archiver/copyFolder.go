package archiver

import (
	"os"
	"path/filepath"
)

func CopyFolder(src, dst string) error {
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if path == src {
			return nil
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			err = os.WriteFile(dstPath, data, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err = copyFolder(path, dstPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
