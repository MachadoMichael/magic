package archiver

import (
	"io"
	"os"
)

func CopyFolder(src, dst string) error {

	println("____", src, dst)
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	println("xxxxxxx", srcFile, dstFile)

	_, err = io.Copy(dstFile, srcFile)
	return err

	// println("on begin: ", src, dst)
	// err := os.MkdirAll(dst, os.ModePerm)
	// if err != nil {
	// 	return err
	// }
	//
	// return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
	// 	path = dst
	// 	if path == src {
	// 		println("path", path, src)
	// 		return nil
	// 	}
	//
	// 	relPath, err := filepath.Rel(src, path)
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	dstPath := filepath.Join(dst, relPath)
	// 	println("-------dstPath", dstPath)
	//
	// 	if !info.IsDir() {
	// 		data, err := os.ReadFile(path)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		err = os.WriteFile(dstPath, data, os.ModePerm)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		err = CopyFolder(path, dstPath)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	//
	// 	return nil
	// })
}
