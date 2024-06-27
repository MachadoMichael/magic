package args

import "github.com/MachadoMichael/magic/pkg/memorizer"

func four(action, alias, path, parameter string) {
	if action == "memorize" && alias != "" && path != "" && parameter == "--file" {
		memorizer.MemorizeTemplate(alias, path, parameter)
	}
}
