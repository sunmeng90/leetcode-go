package main

/*
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:
		1 -> A
		2 -> B
		3 -> C
		...
		26 -> Z
		27 -> AA
		28 -> AB
		...

Example 1:
	Input: 1
	Output: "A"

Example 2:
	Input: 28
	Output: "AB"

Example 3:
	Input: 701
	Output: "ZY"
*/
func convertToTitle(n int) string {
	if n == 0 {
		return ""
	} else {
		n--
		return convertToTitle(n/26) + string('A'+n%26)
	}
}

func main() {
	println(convertToTitle(1))
	println(convertToTitle(26))
	println(convertToTitle(695))
	println(convertToTitle(701))
	println(convertToTitle(1766))
	println(convertToTitle(11116))
}
