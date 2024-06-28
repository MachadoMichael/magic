package args

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const documentation = `Magic Application Documentation
Overview
The Magic application provides a set of commands to memorize folders and files, and to construct templates easily. Below are the available commands and their usage.

Usage
The general syntax for using the Magic application is:

magic <command>

Commands
Memorize a Folder
To memorize a folder, use the following command:

magic memorize <alias> <path>
<alias>: A short name for the memorized path.
<path>: The full path to the folder you want to memorize.

Memorize a File
To memorize a file, use the following command:

magic memorize <alias> <path> --file
<alias>: A short name for the memorized path.
<path>: The full path to the file you want to memorize.
--file: Specifies that the path is to a file.

Construct a Template
To construct a template using a memorized alias, use the following command:

magic build <alias>
<alias>: The alias of the memorized path.

Check Version
To check the version of the Magic application, use one of the following commands:

magic -version
or
magic -v
`

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
		println("magic version: ", version)
		return
	} else if arg == "-help" || arg == "-h" {
		println(documentation)

		return
	} else {
		log.Fatal("No valid arguments were provided. Please provide some valid argment, if look for help -help or -h.")
	}
}
