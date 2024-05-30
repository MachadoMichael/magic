package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MachadoMichael/magic/pkg/handler"
)

var version float32 = 0.1

func main() {
	numberOfArgs := len(os.Args) - 1

	if numberOfArgs == 0 {
		log.Fatal("No arguments were provided. Please provide the required arguments.")
		return
	}

	if numberOfArgs == 1 {
		arg1 := os.Args[1]
		oneArgHandler(arg1)
	}

	if numberOfArgs == 3 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		arg3 := os.Args[3]
		fmt.Printf("args: %v", numberOfArgs)
		threeArgsHandler(arg1, arg2, arg3)
	}

}

func oneArgHandler(arg string) {
	if arg == "-version" || arg == "-v" {
		fmt.Printf("magic version: %f", version)
		return
	} else if arg == "-help" || arg == "-h" {
		fmt.Printf("magic version: %f", version) // change to help commands
		return
	} else {
		log.Fatal("No valid arguments were provided. Please provide some valid argment, if look for help -help or -h.")
	}
}

func threeArgsHandler(action, path, alias string) {
	if action == "create" {
		handler.MemorizeTemplate(path, alias)

	}
}
