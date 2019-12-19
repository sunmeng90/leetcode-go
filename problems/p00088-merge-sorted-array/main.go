package main

import (
	"fmt"
	"github.com/sunmeng90/leetcode-go/util"
)

/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

	Input:
	nums1 = [1,2,3,0,0,0], m = 3
	nums2 = [2,5,6],       n = 3

	Output: [1,2,2,3,5,6]
*/
// compare and insert from beginning involves lot of movement in each step
// how about set from the end of longest slice nums1?
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

//
//func insertAt(nums []int, endIdx int, val int, idx int) {
//	nums[endIdx] = nums[endIdx-1] // copy last
//	for i := endIdx - 1; i > idx; i-- {
//		nums[i] = nums[i-1]
//	}
//	nums[idx] = val
//}

//func insertAt(nums []int, val int, idx int) []int {
//	oldLen := len(nums)
//	if idx >= oldLen {
//		nums = append(nums, val)
//	} else {
//		nums = append(nums, nums[oldLen-1]) // copy last
//		for i := oldLen - 1; i > idx; i-- {
//			nums[i] = nums[i-1]
//		}
//		nums[idx] = val
//	}
//	return nums
//	//fmt.Printf("%s\n", util.IntSliceToString(nums))
//}

func main() {
	nums1 := []int{1, 3, 5, 7, 0, 0, 0, 0}
	nums2 := []int{2, 4, 6, 8}
	merge(nums1, 4, nums2, 4)
	//nums2 = insertAt(nums2, 1, 2)
	fmt.Printf("%s\n", util.IntSliceToString(nums1))

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Printf("%s\n", util.IntSliceToString(nums1))

	nums1 = []int{4, 5, 6, 0, 0, 0}
	nums2 = []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
	fmt.Printf("%s\n", util.IntSliceToString(nums1))

}
