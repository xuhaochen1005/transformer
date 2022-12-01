package utils

func Eext(path string) string {
	var fileExt string
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			break
		}
		if path[i] == '.' {
			fileExt = path[i:]
		}
	}
	return fileExt
}
