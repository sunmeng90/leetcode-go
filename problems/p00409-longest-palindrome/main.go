package main

import "sort"

/*
Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be
built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.

*/
func longestPalindrome(s string) int {
	bytes := []byte(s)
	count := 0
	sort.Sort(byteSlice(bytes))
	for i := 0; i < len(bytes)-1; i++ {
		if bytes[i] == bytes[i+1] {
			count++
			i++
		}
	}
	if len(bytes) > count*2 {
		return count*2 + 1
	}
	return count * 2
}

type byteSlice []byte

func (bs byteSlice) Len() int           { return len(bs) }
func (bs byteSlice) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs byteSlice) Less(i, j int) bool { return bs[i] < bs[j] }

func main() {
	bytes := []byte("abccccdd")
	sort.Sort(byteSlice(bytes))
	println(string(bytes))
}
