package main

import (
	"fmt"
	"strconv"
)

/*
Given an array of characters, compress it in-place.

The length after compression must always be smaller than or equal to the original array.

Every element of the array should be a character (not int) of length 1.

After you are done modifying the input array in-place, return the new length of the array.


Follow up:
Could you solve it using only O(1) extra space?


Example 1:

Input:
["a","a","b","b","c","c","c"]

Output:
Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]

Explanation:
"aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".


Example 2:

Input:
["a"]

Output:
Return 1, and the first 1 characters of the input array should be: ["a"]

Explanation:
Nothing is replaced.


Example 3:

Input:
["a","b","b","b","b","b","b","b","b","b","b","b","b"]

Output:
Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].

Explanation:
Since the character "a" does not repeat, it is not compressed. "bbbbbbbbbbbb" is replaced by "b12".
Notice each digit has it's own entry in the array.


Note:

All characters have an ASCII value in [35, 126].
1 <= len(chars) <= 1000.

*/
func compress(chars []byte) int {
	charStartIdx := 0 // the index of first occurrence char
	numEndIdx := 0    // the index of last digit of count
	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[charStartIdx] {
			if charStartIdx == numEndIdx+1 && chars[charStartIdx] != chars[numEndIdx] {
				numEndIdx++
			}
			count := i - charStartIdx
			if count > 1 {
				countBytes := getDigits(count)
				for j := 0; j < len(countBytes); j++ {
					numEndIdx++
					chars[numEndIdx] = countBytes[j]
				}
			}
			charStartIdx = i

			if charStartIdx >= numEndIdx+1 {
				numEndIdx++
				chars[numEndIdx] = chars[charStartIdx]
			}
		}
	}
	if charStartIdx < len(chars)-1 {
		if charStartIdx == numEndIdx+1 && chars[charStartIdx] != chars[numEndIdx] {
			numEndIdx++
		}
		count := len(chars) - charStartIdx
		countBytes := getDigits(count)
		for j := 0; j < len(countBytes); j++ {
			numEndIdx++
			chars[numEndIdx] = countBytes[j]
		}
	} else if charStartIdx == len(chars) && charStartIdx == numEndIdx+1 {
		return len(chars)
	}
	println(strconv.Itoa(numEndIdx) + " -> " + string(chars) + " -> " + string(chars[:numEndIdx+1]))
	return numEndIdx + 1
}

//
func getDigits(num int) []byte {
	return []byte(fmt.Sprintf("%d", num))
}

func main() {
	compress([]byte("abc"))

	compress([]byte("a"))
	compress([]byte("aa"))
	compress([]byte("aaa"))
	compress([]byte("aaab"))
	compress([]byte("aaabb"))
	compress([]byte("aaabbc"))
	compress([]byte("aaabbcdefg"))
	compress([]byte("aaaaaaaaaaaab"))
	compress([]byte("abc"))
	compress([]byte("abbc"))
	compress([]byte("abb"))
	compress([]byte("abbb"))
	compress([]byte("abbbbbbbbbb"))
	compress([]byte("aabbccc"))
	compress([]byte("aaabbaa"))
}
