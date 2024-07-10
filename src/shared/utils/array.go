package utils

func IsContains[T comparable](arr []T, item T) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}