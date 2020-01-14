package main

/*
Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.

Please note that the string does not contain any non-printable characters.

Example:

Input: "Hello, my name is John"
Output: 5
*/
func countSegments(s string) int {
	count := 0
	wBegin := false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if !wBegin {
				wBegin = true
			}
		} else {
			if wBegin {
				count++
				wBegin = false
			}
		}
	}
	if wBegin {
		count++
	}
	return count
}

func countSegments2(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && (i == 0 || s[i-1] == ' ') { // a word should begin with a blank, or this is the first word
			count++
		}
	}
	return count
}

func main() {
	println(countSegments(""))
	println(countSegments(" "))
	println(countSegments(" 1"))
	println(countSegments(" 1 "))
	println(countSegments(" 1 2"))
	println(countSegments(" 1, 2"))
	println(countSegments(" 1, 2,"))
	println(countSegments(" 1, 2, "))
}
