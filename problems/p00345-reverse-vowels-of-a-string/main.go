package main

import (
	"strings"
)

// scan from beginning and end, and stop and mark if any vowels found, if two found, then swap and reset mark and continue
func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	bytes := []byte(s)
	pair := 0 // bit mask for match vowels from start and end: 00 no match, 01 match from end, 10 match from beginning, 11 two match
	for i, j := 0, len(bytes)-1; i < j; {
		if !strings.Contains(vowels, string(bytes[i])) {
			if pair != 2 { // only if there is no match from start or the matched one is already swapped, then proceed
				i++
			}
		} else {
			pair |= 2 // match from start
			if pair == 3 {
				pair = 0
				bytes[i], bytes[j] = bytes[j], bytes[i]
				i++
				j--
			}
		}
		if !strings.Contains(vowels, string(bytes[j])) {
			if pair != 1 {
				j--
			}
		} else {
			pair |= 1      // match from end
			if pair == 3 { // two match
				pair = 0                                // reset
				bytes[i], bytes[j] = bytes[j], bytes[i] //swap
				i++                                     // proceed
				j--
			}
		}

	}
	return string(bytes)
}

func reverseVowels2(s string) string {
	vowels := "aeiouAEIOU"
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		for ; i < j; i++ {
			if strings.Contains(vowels, string(bytes[i])) {
				break
			}
		}
		for ; i < j; j-- {
			if strings.Contains(vowels, string(bytes[j])) {
				break
			}
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]

	}
	return string(bytes)
}

func main() {
	println(string("abc"[0]))
	println(reverseVowels("hello"))
	println(reverseVowels("leetcode"))
	println(reverseVowels("a.b,."))
}
