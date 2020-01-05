package main

import "strings"

/*
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

Input: pattern = "abba", str = "dog cat cat dog"
Output: true
Example 2:

Input:pattern = "abba", str = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
Example 4:

Input: pattern = "abba", str = "dog dog dog dog"
Output: false
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.
*/

func wordPattern(pattern string, str string) bool {
	charIdxMap := make(map[byte]int, 0)
	strIdxMap := make(map[string]int, 0)
	strSlice := strings.Split(str, " ")
	if len(strSlice) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if charIdxMap[pattern[i]] != strIdxMap[strSlice[i]] {
			return false
		}
		charIdxMap[pattern[i]] = i + 1
		strIdxMap[strSlice[i]] = i + 1
	}
	return true
}

func main() {
	println(wordPattern("abba", "dog cat cat dog"))
	println(wordPattern("abba", "dog dog dog dog"))
}
