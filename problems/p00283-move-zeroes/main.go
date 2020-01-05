package main

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the
non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

// fill the zero with non-zero successor, and then fill zero to the end
func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	countZero := 0
	preNonZero := -1 // keep track the pre non zero index
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			preNonZero++
			nums[preNonZero] = nums[i]
		} else {
			countZero++
		}
	}

	//for i := 0; i < countZero; i++ {
	//	nums[len(nums)-1-i] = 0
	//}
	for i := preNonZero + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}
