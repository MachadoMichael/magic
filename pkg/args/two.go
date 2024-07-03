package args

import (
	"log"

	"github.com/MachadoMichael/magic/pkg/memorizer"
)

func two(action, alias string) {
	if action == "delete" || action == "del" {
		memorizer.DeleteTemplate(alias)
	} else {
		log.Fatalln("Error when we try delete alias.")
	}
}
