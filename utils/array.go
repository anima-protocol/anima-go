package utils

func InArray(str string, array []string) bool {
	for _, item := range array {
		if item == str {
			return true
		}
	}
	return false
}
