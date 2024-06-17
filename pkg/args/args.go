package args

func ReadArgs(args []string) {
	switch len(args) {
	case 1:
		one(args[1])
	}
}
