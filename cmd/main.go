package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MachadoMichael/magic/pkg/args"
	"github.com/MachadoMichael/magic/pkg/mapper"
	"github.com/MachadoMichael/magic/pkg/memorizer"
)

var version float32 = 0.3

func main() {
	mapper.InitMapping()

	numberOfArgs := len(os.Args) - 1
	if numberOfArgs == 0 {
		log.Fatal("No arguments were provided. Please provide the required arguments.")
		return
	}

	if numberOfArgs == 1 {
		args.ReadArgs(os.Args)
	}

	if numberOfArgs == 3 {
		arg1, arg2, arg3 := os.Args[1], os.Args[2], os.Args[3]
		fmt.Println("args: ", numberOfArgs)
		threeArgsHandler(arg1, arg2, arg3)
	}

}

func threeArgsHandler(action, alias, path string) {
	if action == "create" && alias != "" && path != "" {
		memorizer.MemorizeTemplate(alias, path)
	}

}
