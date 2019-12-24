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

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
