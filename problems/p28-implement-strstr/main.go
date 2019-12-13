package main

func main() {
	println(strStr("hello", "ll"))
	println(strStr("heoll", "ll"))
}

func strStr(haystack string, needle string) int {
	hl, nl := len(haystack), len(needle)
	if nl == 0 {
		return 0
	}
	if nl > hl {
		return -1
	}
	for i := 0; i < (hl - nl + 1); i++ {
		if string(haystack[i:i+nl]) == needle {
			return i
		}
	}
	return -1
}
