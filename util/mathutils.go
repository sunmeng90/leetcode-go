package util

func Min(m int, n int) int {
	if m <= n {
		return m
	}
	return n
}

func Max(m int, n int) int {
	if m <= n {
		return n
	}
	return m
}
