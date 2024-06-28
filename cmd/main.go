package main

import (
	"log"
	"os"

	"github.com/MachadoMichael/magic/infra"
	"github.com/MachadoMichael/magic/pkg/args"
	"github.com/MachadoMichael/magic/pkg/mapper"
	"github.com/MachadoMichael/magic/pkg/memorizer"
)

func main() {
	infra.Init()
	mapper.Init()
	memorizer.Init()

	numberOfArgs := len(os.Args) - 1
	if numberOfArgs == 0 {
		log.Fatal("No arguments were provided. Please provide the required arguments.")
		return
	}
	args.ReadArgs(os.Args)

}
