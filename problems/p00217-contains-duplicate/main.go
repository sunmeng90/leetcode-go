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
	println(containsDuplicate([]int{1, 2, 3, 1}))
	println(containsDuplicate([]int{1, 0, 1, 1}))
	println(containsDuplicate([]int{1, 2, 3, 1, 2, 3}))
	println(containsDuplicate2([]int{1, 2, 3, 1, 2, 3}))
	println(containsDuplicate2([]int{3, 1}))
}

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]int)
	for i, v := range nums {
		_, ok := tmp[v]
		if !ok {
			tmp[v] = i
		} else {
			return true
		}
	}
	return false
}

//https://leetcode.com/problems/contains-duplicate/discuss/60937/3ms-Java-Solution-with-Bit-Manipulation
// only work for non-negative value, range from 0 to 15000*8-1
//the multiples of 8 (8的倍数)
// For element at i, start from the Kth element away from i and go back
func containsDuplicate2(nums []int) bool {
	bitMap := make([]byte, 15000)
	for _, v := range nums {
		segIdx, bitIdx := v/8, byte(v%8) // calculate the segment(byte) index in array, and the bit index of remainder in the segment
		bitIdx = 1 << bitIdx             //move the bit(1) in the segment to the index of remainder
		if (bitMap[segIdx] & bitIdx) != 0 {
			return true
		}
		bitMap[segIdx] |= bitIdx
	}
	return false
}

//TODO use two array one for negative and one for positive numbers to support negative numbers

// build a bit map for all possible input numbers(still limited)
// one bit can represent a number
// how?
// an array of bits that first partitioned into fixed segment
// the segment index of the number is num / segment_size
// the index in the segment of the number is num % segment_size
// [00000000,00000000,00000000,00000000,00000000,00000000,00000000]
// if the segment size is 8, that is a byte
// map for number 8 will be
// [00000000,10000000,00000000,00000000,00000000,00000000,00000000]
// add another map for number 11
// [00000000,10010000,00000000,00000000,00000000,00000000,00000000]
// each segment can represent segment_size numbers, and the number range for that segment will be
// i*segment_size to i*segment_size -1
