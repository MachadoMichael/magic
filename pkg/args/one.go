package args

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func one(arg string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error to read .env file")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatalln("Error to read .env file")
	}

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
