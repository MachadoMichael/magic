package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MachadoMichael/magic/pkg/handler/mapper"
	"github.com/MachadoMichael/magic/pkg/handler/memorizer"
)

var version float32 = 0.1

func main() {
	mapper.InitMapping()

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
		arg1, arg2, arg3 := os.Args[1], os.Args[2], os.Args[3]
		fmt.Println("args: ", numberOfArgs)
		threeArgsHandler(arg1, arg2, arg3)
	}

}

func oneArgHandler(arg string) {
	if arg == "-version" || arg == "-v" {
		fmt.Println("magic version: ", version)
		return
	} else if arg == "-help" || arg == "-h" {
		fmt.Println("magic version: ", version) // change to help commands
		return
	} else {
		log.Fatal("No valid arguments were provided. Please provide some valid argment, if look for help -help or -h.")
	}
}

func threeArgsHandler(action, alias, path string) {
	if action == "create" && alias != "" && path != "" {
		memorizer.MemorizeTemplate(alias, path)
	}

}
