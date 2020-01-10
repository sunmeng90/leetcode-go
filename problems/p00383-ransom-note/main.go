package main

/*
Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function
that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
*/
func canConstruct(ransomNote string, magazine string) bool {
	charMap := make([]int, 26)
	for i := 0; i < len(ransomNote); i++ {
		charMap[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		charMap[magazine[i]-'a']--
	}
	for i := 0; i < len(charMap); i++ {
		if charMap[i] > 0 {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	charCountMap := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		charCountMap[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		charCountMap[ransomNote[i]-'a']--
		if charCountMap[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
}
