package args

import (
	"github.com/MachadoMichael/magic/pkg/memorizer"
)

func three(action, alias, path string) {
	if action == "memorize" && alias != "" && path != "" {
		memorizer.MemorizeTemplate(alias, path, "")
	}

}
