package main

import "github.com/sunmeng90/leetcode-go/util"

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

	Input: [1,2,3,1]
	Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

	Input: [2,7,9,3,1]
	Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/
// https://leetcode.com/problems/house-robber/discuss/156523/From-good-to-great.-How-to-approach-most-of-DP-problems.

func rob(nums []int) int {
	memo := make([]int, len(nums)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return robAt(nums, len(nums)-1, memo)
}

func robAt(nums []int, i int, memo []int) int {
	if i < 0 {
		return 0
	}
	if memo[i] >= 0 {
		return memo[i]
	}
	memo[i] = util.Max(robAt(nums, i-2, memo)+nums[i], robAt(nums, i-1, memo))
	return memo[i]
}

// iterative way
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums)+1)
	memo[0], memo[1] = 0, nums[0]
	for i := 1; i < len(nums); i++ {
		memo[i+1] = util.Max(memo[i], memo[i-1]+nums[i])
	}
	return memo[len(nums)]
}

func main() {
	println(rob2([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81,
		90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169,
		241, 202, 144, 240}))
}
