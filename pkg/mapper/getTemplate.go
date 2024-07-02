package mapper

func GetTemplate(alias string) (InfoMap, error) {
	return magicMap[alias], nil
}
