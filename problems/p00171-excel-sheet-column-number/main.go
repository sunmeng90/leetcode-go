package main

/*
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
Example 1:

	Input: "A"
	Output: 1
Example 2:

	Input: "AB"
	Output: 28
Example 3:

	Input: "ZY"
	Output: 701
*/
func titleToNumber(s string) int {
	result := 0
	radix := 1
	for i := len(s) - 1; i >= 0; i-- {
		result += (int(s[i]) - 64) * radix
		radix *= 26
	}
	return result
}

func titleToNumber2(s string) int {
	result := 0
	for _, v := range s {
		result = result*26 + (int(v) - 64)
	}
	return result
}

func main() {
	println(titleToNumber("A"))
	println(titleToNumber("Z"))
	println(titleToNumber("AA"))
	println(titleToNumber("ABC"))
	println(titleToNumber("ABCD"))
	println(titleToNumber("ABCDE"))

	println(titleToNumber2("A"))
	println(titleToNumber2("Z"))
	println(titleToNumber2("AA"))
	println(titleToNumber2("ABC"))
	println(titleToNumber2("ABCD"))
	println(titleToNumber2("ABCDE"))
}
