package main

/*
Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the
maximum number. The time complexity must be in O(n).

Example 1:
Input: [3, 2, 1]

Output: 1

Explanation: The third maximum is 1.
Example 2:
Input: [1, 2]

Output: 2

Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
Example 3:
Input: [2, 2, 3, 1]

Output: 1

Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.
*/

// if a large number is arriving, then should move current numbers forwardly(from first -> third)
func thirdMax(nums []int) int {
	firstMax, secondMax, thirdMax := 0, 0, 0
	firstInit, secondInit, thirdInit := false, false, false
	for i := 0; i < len(nums); i++ {
		if (firstInit && nums[i] == firstMax) || (secondInit && nums[i] == secondMax) || (thirdInit && nums[i] == thirdMax) {
			continue
		}
		if nums[i] > firstMax || !firstInit {
			if secondInit {
				thirdMax = secondMax
				thirdInit = true
			}
			if firstInit {
				secondMax = firstMax
				secondInit = true
			}
			firstMax = nums[i]
			firstInit = true
		} else if nums[i] < firstMax && (nums[i] > secondMax || !secondInit) {
			if secondInit {
				thirdMax = secondMax
				thirdInit = true
			} else {
				secondInit = true
			}
			secondMax = nums[i]
		} else if nums[i] < secondMax && (nums[i] > thirdMax || !thirdInit) {
			thirdMax = nums[i]
			thirdInit = true
		}
	}
	if thirdInit {
		return thirdMax
	}
	return firstMax
}

func main() {
	println(thirdMax([]int{3, 2, 1}))
	println(thirdMax([]int{1, 2, -2147483648}))
	println(thirdMax([]int{-2147483648, 1, 1}))
}
