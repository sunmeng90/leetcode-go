package main

import "fmt"

/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

	Input: [1,2,3]
	Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

	Input: [4,3,2,1]
	Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/
func plusOne(digits []int) []int {
	overflow := true // image an overflow beyond the end of last digit
	endIndex := len(digits) - 1
	for i := endIndex; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			overflow = false
			break
		}
	}
	if overflow { // if overflow to most significant order
		digits = append(digits, digits[endIndex])
		for endIndex > 0 {
			digits[endIndex] = digits[endIndex-1]
			endIndex--
		}
		digits[0] = 1
	}
	return digits
}

func main() {
	fmt.Printf("%v", plusOne([]int{1, 2, 3}))
	fmt.Printf("%v", plusOne([]int{9}))
	fmt.Printf("%v", plusOne([]int{0}))
	fmt.Printf("%v", plusOne([]int{99}))
}
