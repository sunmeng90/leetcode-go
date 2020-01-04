package main

/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?

*/
// build char frequency map from first string, and decrease it by occurrence in second string
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charFreq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		charFreq[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		charFreq[t[i]-'a']--
		if charFreq[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charFreq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		charFreq[s[i]-'a']++
		charFreq[t[i]-'a']--
	}
	for _, v := range charFreq {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	println(isAnagram("a", "b"))
}
