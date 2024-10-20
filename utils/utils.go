package utils

import "strconv"

func StringToUint(s string) uint {
	result, _ := strconv.ParseUint(s, 10, 64)
	return uint(result)
}

func StringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
