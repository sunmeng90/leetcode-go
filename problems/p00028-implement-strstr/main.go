package main

import "fmt"

func strStr(haystack string, needle string) int {
	hl, nl := len(haystack), len(needle)
	if nl == 0 {
		return 0
	}
	if nl > hl {
		return -1
	}
	for i := 0; i < (hl - nl + 1); i++ {
		if string(haystack[i:i+nl]) == needle {
			return i
		}
	}
	return -1
}

//TODO: KMP (Knuth Morris Pratt) Pattern Searching
// text
// pat
// pre-processing: find the all the sub-pattern(not the full sub-string in pattern) in pattern that is both a prefix and suffix
//				for sub-string of the pattern
// matching:
// when comparing the pattern in text, if there are n character in text matches from the beginning of pattern and next
// character in text not match, the n matched characters(which is suffix for sub-string in the pattern) will be the prefix
// of pattern and no need to compared again, we can continue to compare with the n+1 characters in the pattern from the
// beginning(prefix) of pattern
// =====================================================================================================================
// simply speaking, when the suffix is found and the n+1 doesn't match, then the characters in suffix which is also prefix
// for the pattern can be skipped,continue with the next character in the pattern
// text: abcxabcdabxabcdabcdabcy
// pattern: abcdabcy
// there is a sub-string 'ab' which is both a prefix and suffix for sub-string 'abcdab' in pattern
// there is a sub-string 'abc' which is both a prefix and suffix for sub-string 'abcdabc' in pattern
// reference: Longest Prefix Suffix

//	Examples of lps[] construction:
//	For the pattern “AAAA”,
//	lps[] is [(0), 1, 2, 3]
//	(0): for sub-pattern 'A', the longest prefix and suffix is 'A', but full length of sub-pattern is not allowed, so it is 0,
//  same explanation for 1, 2, 3

//	For the pattern “ABCDE”,
//	lps[] is [0, 0, 0, 0, 0]
//
//	For the pattern “AABAACAABAA”,
//	lps[] is [0, 1, 0, 1, (2), 0, 1, 2, 3, 4, 5]
//  (2): for sub-pattern AABAA, the longest prefix and suffix is AA, and the length is xxxix is 2
//
//	For the pattern “AAACAAAAAC”,
//	lps[] is [0, 1, 2, 0, 1, 2, 3, 3, 3, 4]
//
//	For the pattern “AAABAAA”,
//	lps[] is [0, 1, 2, 0, 1, 2, 3]

// when q = 0, then k = 0
// when q > 0, and pattern[q] == pattern[k], then k++ and result[q] = k
// when q > 0, and pattern[q] != pattern[k] and k > 0, then go back for previous result[k-1] and compare again,
//				until k = 0, that means not a prefix and suffix, so result[q]=0, then we move on the next character in
//				the pattern
// maps q to the length of the longest prefix of Pattern that is a suffix of Pattern[0 . . q]
func computePrefix(pattern string) []int {
	m := len(pattern) // pattern length
	result := make([]int, m)
	result[0] = 0
	k := 0 // previous longest prefix suffix
	for q := 1; q < m; {
		if pattern[q] == pattern[k] {
			k++
		} else if k > 0 { // if the previous can't not continue(pattern[q] != pattern[k])
			// ABA(B) :[0 0 1 2]
			// ABCAB(C) : [0 0 0 1 2 3]
			// really tricky, can't understand
			k = result[k-1] // this is NOT go back one from index q in result
			continue        // q not change, because we don't find the prefix yet for substring pattern[0:q(inclusive)], go back one previous prefix further
		}
		result[q] = k
		q++
	}
	return result
}

func kmp(text, pattern string) {
	prefix := computePrefix(pattern)
	for i, j := 0, 0; i < len(text); { // i: index for text, j index for pattern
		if pattern[j] == text[i] {
			i++
			j++
		}
		if j == len(pattern) { // match the whole pattern
			fmt.Printf("Found pattern at index: %d\n", i-j)
			j = prefix[j-1] //TODO figure it out
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = prefix[j-1]
			} else {
				i++
			}
		}
	}
}

func main() {
	println(strStr("hello", "ll"))
	println(strStr("heoll", "ll"))
	//fmt.Printf("%v\n", computePrefix("heoll"))
	fmt.Printf("%v\n", computePrefix("AAAA"))        //[(0), 1, 2, 3]
	fmt.Printf("%v\n", computePrefix("ABCDE"))       //[0, 0, 0, 0, 0]
	fmt.Printf("%v\n", computePrefix("AABAACAABAA")) //[0, 1, 0, 1, (2), 0, 1, 2, 3, 4, 5]
	fmt.Printf("%v\n", computePrefix("ABAB"))        //[0 0 1 2]
	fmt.Printf("%v\n", computePrefix("ABABAA"))      //[0 0 1 2 3 1]
	fmt.Printf("%v\n", computePrefix("ABCABC"))      //[0 0 0 1 2 3]
	kmp("abcabcabcdefgabc", "abc")
}
