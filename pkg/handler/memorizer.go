package handler

import (
	"fmt"
)

var destination string = "../../templates/"

func MemorizeTemplate(alias, path string) {
	fmt.Printf("path: %s", path)

	initMapping()
	setNewAlias(alias, "caminhos")

	// err := os.WriteFile(path+alias, []byte("hello, Gophers!"), 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//

	// files, err := os.ReadDir(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// newProjectPath := filepath.Join(alias)
	// err = os.Mkdir(newProjectPath, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, file := range files {
	// 	srcFile := filepath.Join(path, file.Name())
	//
	// 	err = copyFile(srcFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}

func copyFile(src string) error {
	fmt.Printf("%v %v ", src, destination)
	return nil
}
