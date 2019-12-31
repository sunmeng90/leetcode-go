package main

/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:
	Input: [3,2,3]
	Output: 3

Example 2:
	Input: [2,2,1,1,1,2,2]
	Output: 2
*/
/*
Boyer-Moore Voting Algorithm
Compare numbers in array, if two number are different then ignore them from result, if they are equal, keep it, finally
if there is a majority, it will be the only numbers in the end
*/
func majorityElement(nums []int) int {
	res := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if v == res {
			count++
		} else {
			count--
		}
	}
	return res
}
