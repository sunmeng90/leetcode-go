package main

import "fmt"

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.
*/
func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, start, end int) int {
	p := (start + end) / 2
	if nums[p] == target {
		return p // found
	}
	if start == end {
		if target < nums[start] { // if the target is smaller first, then choose the start position
			return start
		} else { // if the target is large than the second, then take position after the second(end+1)
			return end + 1
		}
	} else if (end - start) == 1 {
		if target < nums[start] { // if the target is smaller first, then choose the start position
			return start
		} else if target <= nums[end] { // if the target is large than first and smaller than or equal the second, then choose the end position
			return end
		} else { // if the target is large than the second, then take position after the second(end+1)
			return end + 1
		}
	}

	if target < nums[p] {
		return binarySearch(nums, target, start, p-1)
	} else {
		return binarySearch(nums, target, p+1, end)
	}
}

func searchInsert2(nums []int, target int) int {
	return binarySearch2(nums, target, 0, len(nums)-1)
}

func binarySearch2(nums []int, target int, start, end int) int {
	p := (start + end) / 2
	if nums[p] == target {
		return p // found
	}
	if (end - start) <= 1 {
		if target < nums[start] { // if the target is smaller first, then choose the start position
			return start
		} else if target <= nums[end] { //[only there are two value] if the target is large than first and smaller than or equal the second, then choose the end position
			return end
		} else { // if the target is large than the second, then take position after the second(end+1)
			return end + 1
		}
	}

	if target < nums[p] {
		return binarySearch(nums, target, start, p-1)
	} else {
		return binarySearch(nums, target, p+1, end)
	}
}

// no recursion
func searchInsert3(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2 // choose the lower one in the middle if the count of numbers is even
		if target <= nums[mid] { // the target is in the middle or before it
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

//https://www.golangprograms.com/functions-mess-recursive-anonymous-function-in-golang.html
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))      //2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))      //0
	fmt.Println(searchInsert([]int{3, 5, 7, 9, 10}, 8))  //3
	fmt.Println(searchInsert([]int{3, 5, 7, 9, 10}, 11)) //5
	fmt.Println(searchInsert([]int{1, 3}, 2))            //1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))      //1
	fmt.Println(searchInsert([]int{1, 3}, 3))            //1
	fmt.Println("================")
	fmt.Println(searchInsert3([]int{1, 3, 5, 6}, 5))      //2
	fmt.Println(searchInsert3([]int{1, 3, 5, 6}, 0))      //0
	fmt.Println(searchInsert3([]int{3, 5, 7, 9, 10}, 8))  //3
	fmt.Println(searchInsert3([]int{3, 5, 7, 9, 10}, 11)) //5
	fmt.Println(searchInsert3([]int{1, 3}, 2))            //1
	fmt.Println(searchInsert3([]int{1, 3, 5, 6}, 2))      //1
	fmt.Println(searchInsert3([]int{1, 3}, 3))            //1
}
