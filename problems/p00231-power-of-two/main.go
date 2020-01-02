package main

/*
Given an integer, write a function to determine if it is a power of two.

Example 1:
	Input: 1
	Output: true
Explanation: 20 = 1

Example 2:
	Input: 16
	Output: true
Explanation: 24 = 16

Example 3:
	Input: 218
	Output: false
*/
func isPowerOfTwo(n int) bool {
	count := 0 // count of bit '1'
	for n > 0 {
		if n&1 == 1 {
			count++
			if count > 1 {
				break
			}
		}
		n >>= 1
	}
	return count == 1
}
