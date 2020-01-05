package main

/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
Note:
You may assume that the array does not change.
There are many calls to sumRange function.
*/

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for ; i <= j; i++ {
		sum += this.nums[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
