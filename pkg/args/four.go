package args

import "github.com/MachadoMichael/magic/pkg/memorizer"

func four(action, alias, path, parameter string) {
	if action == "save" && alias != "" && path != "" && parameter == "--file" {
		memorizer.SaveTemplate(alias, path, parameter)
	}
}
