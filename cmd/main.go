package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MachadoMichael/magic/pkg/args"
	"github.com/MachadoMichael/magic/pkg/mapper"
)

var version float32 = 0.3

func main() {
	mapper.InitMapping()

	numberOfArgs := len(os.Args) - 1
	if numberOfArgs == 0 {
		log.Fatal("No arguments were provided. Please provide the required arguments.")
		return
	}

	args.ReadArgs(os.Args)
	fmt.Println("args: ", numberOfArgs)

}
