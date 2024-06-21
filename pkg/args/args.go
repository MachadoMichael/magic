package args

func ReadArgs(args []string) {
	switch len(args) {
	case 2:
		one(args[1])
	case 4:
		three(args[1], args[2], args[3])
	case 5:
		four(args[1], args[2], args[3], args[4])
	}
}
