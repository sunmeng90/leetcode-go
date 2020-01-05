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
	nums     []int
	sumCache []int // ith element would be the sum of element from (0 to i)
}

func Constructor(nums []int) NumArray {
	caches := make([]int, len(nums)+1) // real sum start from caches[1]
	for i := 0; i < len(nums); i++ {
		caches[i+1] = caches[i] + nums[i]
	}
	return NumArray{nums: nums, sumCache: caches}
}

// cache result of sum(1, k), and calculate the range sum by sum[j+1] - sum[i]
func (this *NumArray) SumRange(i int, j int) int {
	return this.sumCache[j+1] - this.sumCache[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
