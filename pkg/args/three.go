package args

import (
	"github.com/MachadoMichael/magic/pkg/memorizer"
)

func three(action, alias, path string) {
	if action == "create" && alias != "" && path != "" {
		memorizer.MemorizeTemplate(alias, path)
	}

}
