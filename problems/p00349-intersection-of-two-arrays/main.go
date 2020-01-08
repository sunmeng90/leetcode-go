package main

/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.

*/
func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]bool, 0)
	nums2Map := make(map[int]bool, 0)
	result := make([]int, 0)
	for _, v := range nums1 {
		nums1Map[v] = true
	}
	for _, v := range nums2 {
		nums2Map[v] = true
	}
	for k, _ := range nums2Map {
		_, ok := nums1Map[k]
		if ok {
			result = append(result, k)
		}
	}
	return result
}
