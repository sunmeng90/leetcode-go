package main

import "sort"

/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Note:

Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the
memory at once?
*/

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		n1, n2 := nums1[i], nums2[j]
		if n1 == n2 {
			result = append(result, n1)
			i++
			j++
		}
		if n1 < n2 {
			i++
		} else if n1 > n2 {
			j++
		}
	}
	return result
}
