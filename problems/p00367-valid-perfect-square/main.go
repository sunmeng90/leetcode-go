package main

/*
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Output: true
Example 2:

Input: 14
Output: false
*/
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i := num / 2; i > 1; i-- {
		if num == i*i {
			return true
		}
	}
	return false
}
