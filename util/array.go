package util

func Contains[T comparable](array []T, el T) bool {
	for _, i := range array {
		if i == el {
			return true
		}
	}
	return false
}
