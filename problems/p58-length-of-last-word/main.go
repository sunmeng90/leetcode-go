package main

func main() {
	println(lengthOfLastWord("ab c d ef 123abc"))
}

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

	Input: "Hello World"
	Output: 5

*/
func lengthOfLastWord(s string) int {
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && sum > 0 {
			break
		} else if s[i] != ' ' {
			sum++
		}
	}
	return sum
}
