package main

import (
	"fmt"
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
				numEndIdx += addCount(&chars, count, numEndIdx+1)
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
		numEndIdx += addCount(&chars, count, numEndIdx+1)
	} else if charStartIdx == len(chars) && charStartIdx == numEndIdx+1 {
		return len(chars)
	}
	//println(strconv.Itoa(numEndIdx) + " -> " + string(chars) + " -> " + string(chars[:numEndIdx+1]))
	return numEndIdx + 1
}

// convert to number to digits
func getDigits(num int) []byte {
	return []byte(fmt.Sprintf("%d", num))
}

func addCount(bytes *[]byte, count int, start int) int {
	bitCount := 0
	countBak := count
	for count > 0 {
		bitCount++
		count /= 10
	}
	end := start + bitCount - 1
	for i := bitCount; i > 0; i-- {
		(*bytes)[end] = '0' + uint8(countBak%10)
		countBak /= 10
		end--
	}
	return bitCount
}

func assertEquals(src, expected, actual string) {
	if expected != actual {
		panic(fmt.Sprintf("Error: src: %s, expected: %s, actual: %s\n", src, expected, actual))
	} else {
		fmt.Printf("%s -> %s\n", src, actual)
	}
}

func main() {
	src := "a"
	srcBytes := []byte(src)
	assertEquals(src, "a", string(srcBytes[:compress(srcBytes)]))
	src = "aa"
	srcBytes = []byte(src)
	assertEquals(src, "a2", string(srcBytes[:compress(srcBytes)]))
	src = "aaa"
	srcBytes = []byte(src)
	assertEquals(src, "a3", string(srcBytes[:compress(srcBytes)]))

	src = "aaab"
	srcBytes = []byte(src)
	assertEquals(src, "a3b", string(srcBytes[:compress(srcBytes)]))

	src = "aaabb"
	srcBytes = []byte(src)
	assertEquals(src, "a3b2", string(srcBytes[:compress(srcBytes)]))

	src = "aaabbc"
	srcBytes = []byte(src)
	assertEquals(src, "a3b2c", string(srcBytes[:compress(srcBytes)]))

	src = "aaabbcdefg"
	srcBytes = []byte(src)
	assertEquals(src, "a3b2cdefg", string(srcBytes[:compress(srcBytes)]))

	src = "aaaaaaaaaaaab"
	srcBytes = []byte(src)
	assertEquals(src, "a12b", string(srcBytes[:compress(srcBytes)]))

	src = "abc"
	srcBytes = []byte(src)
	assertEquals(src, "abc", string(srcBytes[:compress(srcBytes)]))

	src = "abbc"
	srcBytes = []byte(src)
	assertEquals(src, "ab2c", string(srcBytes[:compress(srcBytes)]))

	src = "abb"
	srcBytes = []byte(src)
	assertEquals(src, "ab2", string(srcBytes[:compress(srcBytes)]))

	src = "abbb"
	srcBytes = []byte(src)
	assertEquals(src, "ab3", string(srcBytes[:compress(srcBytes)]))

	src = "abbbbbbbbbb"
	srcBytes = []byte(src)
	assertEquals(src, "ab10", string(srcBytes[:compress(srcBytes)]))

	src = "aabbccc"
	srcBytes = []byte(src)
	assertEquals(src, "a2b2c3", string(srcBytes[:compress(srcBytes)]))

	src = "aaabbaa"
	srcBytes = []byte(src)
	assertEquals(src, "a3b2a2", string(srcBytes[:compress(srcBytes)]))

	src = "a100"
	srcBytes = []byte(src)
	assertEquals(src, "a102", string(srcBytes[:compress(srcBytes)]))

	src = "aaaaaabbbbbbbbbbbbbbbbbbbbbcccccccccccccc"
	srcBytes = []byte(src)
	assertEquals(src, "a6b21c14", string(srcBytes[:compress(srcBytes)]))
}
