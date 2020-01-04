package main

/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two
characters may map to the same character but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true
Note:
You may assume both s and t have the same length.
*/
// build a map of characters from s to t, one character(any position) can only map to the same character in another string
func isIsomorphic(s string, t string) bool {
	charMap := make(map[byte]byte)
	reverseMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		cs, ok := charMap[s[i]] // if char at i is mapped already
		if ok {
			if cs != t[i] {
				return false
			}
		} else {
			cs, ok := reverseMap[t[i]]
			if ok && cs != s[i] { // if char is mapped to another char from t to s
				return false
			}
			charMap[s[i]] = t[i] // add new map entry
			reverseMap[t[i]] = s[i]
		}
	}
	return true
}

// keep update the char index in two string, the counterpart chars in s and t should have the same index
func isIsomorphic2(s string, t string) bool {
	sCharMap := make(map[byte]int, 256)
	tCharMap := make(map[byte]int, 256)
	for i := 0; i < len(s); i++ {
		if sCharMap[s[i]] != tCharMap[t[i]] {
			return false
		}
		sCharMap[s[i]] = i + 1
		tCharMap[t[i]] = i + 1
	}
	return true
}
