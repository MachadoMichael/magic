package args

import (
	"log"

	"github.com/MachadoMichael/magic/pkg/mapper"
)

func two(action, alias string) {
	if action == "delete" || action == "del" {
		mapper.DeleteAlias(alias)
	} else {
		log.Fatalln("Error when we try delete alias.")
	}
}
