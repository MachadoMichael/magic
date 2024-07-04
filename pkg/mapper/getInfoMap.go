package mapper

func GetInfoMap(alias string) (InfoMap, error) {
	return magicMap[alias], nil
}
