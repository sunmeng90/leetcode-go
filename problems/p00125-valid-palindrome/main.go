package main

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

	Input: "A man, a plan, a canal: Panama"
	Output: true
Example 2:

	Input: "race a car"
	Output: false
*/
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		if isAlphaNum(s[i]) && isAlphaNum(s[j]) {
			if isAlphaNumEqualIgnoreCase(s[i], s[j]) {
				i++
				j--
			} else {
				return false
			}
		} else {
			if !isAlphaNum(s[i]) {
				i++
			}
			if !isAlphaNum(s[j]) {
				j--
			}
		}
	}
	return true
}

func isAlphaNum(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func isAlphaNumEqualIgnoreCase(a byte, b byte) bool {
	if a == b {
		return true
	}
	if a <= '9' || b <= '9' {
		return false
	}
	if a > b { // 'P' - '0' = 32 = 'a' - 'A'
		return a-b == 'a'-'A'
	} else {
		return b-a == 'a'-'A'
	}
}

func main() {
	//println('a')
	//println('z')
	//println('A')
	//println('Z')
	//println('0')
	//println('9')
	//println("abc"[0] == 97)
	//println("abc"[0] == 'a')
	//println(isPalindrome("A man, a plan, a canal: Panama"))
	//println(isPalindrome("race a car"))
	println(isPalindrome("0P"))
}
