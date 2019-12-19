package util

import "strconv"

func IntSliceToString(data []int) string {
	result := ""
	if len(data) == 0 {
		return result
	}
	for i := 0; i < len(data); i++ {
		result += strconv.Itoa(data[i]) + ","
	}
	return result[:len(result)-1]
}
