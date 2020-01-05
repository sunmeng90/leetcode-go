package main

/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/

func missingNumber(nums []int) int {
	expectedSum := (len(nums) + 1) * len(nums) / 2
	for i := 0; i < len(nums); i++ {
		expectedSum -= nums[i]
	}
	return expectedSum
}

func missingNumber2(nums []int) int {
	expectedSum := len(nums)
	for i := 0; i < len(nums); i++ {
		expectedSum += i - nums[i] // do the addition and deduction at the same time
	}
	return expectedSum
}

func missingNumber3(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= i
		res ^= nums[i]
	}
	return res
}
