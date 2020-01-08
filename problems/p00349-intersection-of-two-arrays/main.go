package main

import (
	"fmt"
	"sort"
)

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

func intersection2(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			matched := nums1[i]
			result = append(result, matched)
			for ; i < len(nums1) && nums1[i] == matched; i++ { // skip for duplicate one
			}
			for ; j < len(nums2) && nums2[j] == matched; j++ {
			}
		}
	}
	return result
}

func intersection3(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		n1, n2 := nums1[i], nums2[j]
		if n1 == n2 {
			result = append(result, n1)
			for ; i < len(nums1) && nums1[i] == n1; i++ { // skip for duplicate one
			}
			for ; j < len(nums2) && nums2[j] == n1; j++ {
			}
		}
		if n1 < n2 {
			i++
		} else if n1 > n2 {
			j++
		}
	}
	return result
}

// binary search: sort one array, and search element for another arr -> sort the short one
//
func main() {
	fmt.Printf("%v\n", intersection2([]int{1, 2, 2, 1}, []int{2, 2}))
}
