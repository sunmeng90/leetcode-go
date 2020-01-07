package main

import (
	"strings"
)

// scan from beginning and end, and stop and mark if any vowels found, if two found, then swap and reset mark and continue
func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	bytes := []byte(s)
	pair := 0 // bit mask for match vowels from start and end
	for i, j := 0, len(bytes)-1; i < j; {
		if !strings.Contains(vowels, string(bytes[i])) {
			if pair != 2 {
				i++
			}
		} else {
			pair |= 2
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
			pair |= 1
			if pair == 3 {
				pair = 0
				bytes[i], bytes[j] = bytes[j], bytes[i]
				i++
				j--
			}
		}

	}
	return string(bytes)
}

func main() {
	println(string("abc"[0]))
	println(reverseVowels("hello"))
	println(reverseVowels("leetcode"))
	println(reverseVowels("a.b,."))
}
