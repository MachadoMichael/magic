package args

import (
	"github.com/MachadoMichael/magic/pkg/constructor"
	"github.com/MachadoMichael/magic/pkg/memorizer"
)

func three(action, alias, path string) {
	if action == "save" && alias != "" && path != "" {
		memorizer.SaveTemplate(alias, path, "")
	}

	if action == "build" && path != "" {
		constructor.Build(alias, path)
	}

}
