package main

import (
	"fmt"
	"github.com/sunmeng90/leetcode-go/util"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"abc", "ab"}))
	fmt.Println(longestCommonPrefix([]string{"abc", "", "abc"}))
	fmt.Println(longestCommonPrefix([]string{"abc", "ab", "abc"}))
	fmt.Println(longestCommonPrefix([]string{"abc", "ab", "abc", "c"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0] // strs[0][:] generate another string share the same characters
	for _, v := range strs {
		minLen := util.Min(len(prefix), len(v))
		newCommonLength := 0
		for j, commonChar := range prefix[:minLen] { //range over generate a sequence of rune
			if byte(commonChar) != v[j] {
				break
			} else {
				newCommonLength = j + 1
			}
		}
		if newCommonLength == 0 {
			return ""
		}
		prefix = prefix[:newCommonLength]
	}
	return prefix
}

//TODO: divide and conquer
// divide string array until there are two/one elements in the group and then find the common string
// find common string among previous common strings

//TODO: binary search
//The idea is to apply binary search method to find the string with maximum value L, which is common prefix of all of the strings. The algorithm searches space is the interval (0 \ldots minLen)(0…minLen), where minLen is minimum string length and the maximum possible common prefix. Each time search space is divided in two equal parts, one of them is discarded, because it is sure that it doesn't contain the solution. There are two possible cases:
//
//S[1...mid] is not a common string. This means that for each j > i S[1..j] is not a common string and we discard the second half of the search space.
//S[1...mid] is common string. This means that for for each i < j S[1..i] is a common string and we discard the first half of the search space, because we try to find longer common prefix.
