package utils

func GetDirAndFileName(path string) (dir string, fileName string) {
	dirEndIndex := -1
	for i := 0; i < len(path); i++ {
		if string(path[i]) == "/" {
			dirEndIndex = i
		}
	}
	if dirEndIndex == -1 {
		return ".", path
	}
	return path[0:dirEndIndex], path[dirEndIndex+1:]
}
