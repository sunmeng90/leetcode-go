package main

//Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such
//that nums[i] = nums[j] and the absolute difference between i and j is at most k.
//
//Example 1:
//
//Input: nums = [1,2,3,1], k = 3
//Output: true
//Example 2:
//
//Input: nums = [1,0,1,1], k = 1
//Output: true
//Example 3:
//
//Input: nums = [1,2,3,1,2,3], k = 2
//Output: false

func main() {
	println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	println(containsNearbyDuplicate2([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	tmp := make(map[int]int)
	for i, v := range nums {
		oldI, ok := tmp[v]
		if !ok {
			tmp[v] = i
		} else {
			diff := i - oldI
			if diff <= k {
				return true
			} else {
				tmp[v] = i
			}
		}
	}
	return false
}

// For element at i, start from the Kth element away from i and go back
func containsNearbyDuplicate2(nums []int, k int) bool {
	for i, v := range nums {
		tailIdx := i + k
		if tailIdx >= len(nums) {
			tailIdx = len(nums) - 1
		}
		for tailIdx > i {
			if nums[tailIdx] == v {
				return true
			}
			tailIdx--
		}
	}
	return false
}
