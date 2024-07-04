package mapper

func PrintAllInfoMap() error {
	if len(magicMap) == 0 {
		println("There are no maps created")
		return nil
	}

	for alias, data := range magicMap {
		println("Alias: ", alias)
		println("Resource: ", data.Resource)
		println("Path: ", data.Path)
		println("CreateAt: ", data.CreateAt)
		println("Builds: ", data.Builds)
		if data.IsFile {
			println("Type: file")
		} else {
			println("Type: folder")
		}

		println()
	}
	return nil
}
