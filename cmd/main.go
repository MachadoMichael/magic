package main

import (
	"log"
	"os"

	"github.com/MachadoMichael/magic/pkg/args"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

func main() {
	mapper.InitMapping()

	numberOfArgs := len(os.Args) - 1
	if numberOfArgs == 0 {
		log.Fatal("No arguments were provided. Please provide the required arguments.")
		return
	}

	println("args: ", numberOfArgs)
	args.ReadArgs(os.Args)

}
