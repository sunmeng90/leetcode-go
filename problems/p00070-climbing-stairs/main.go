package main

import "fmt"

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

	Input: 2
	Output: 2
	Explanation: There are two ways to climb to the top.
	1. 1 step + 1 step
	2. 2 steps

Example 2:

	Input: 3
	Output: 3
	Explanation: There are three ways to climb to the top.
	1. 1 step + 1 step + 1 step
	2. 1 step + 2 steps
	3. 2 steps + 1 step
*/
/**
time cost is too large
*/
func climbStairs2(n int) int {
	if n < 3 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

/**
dynamic programming, use previous computation result
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dps := make([]int, n+1)
	dps[0] = 0
	dps[1] = 1
	dps[2] = 2
	for i := 3; i <= n; i++ {
		dps[i] = dps[i-1] + dps[i-2]
	}
	return dps[n]
}

/*
Explanation
	i=0: 0
	i=1: 1
	i=2: 2
when i=n
	we can only reach to i-1(one step), and i-2(two step), so the result would be the ways by f(i-1) + f(i-2)

*/

func main() {
	fmt.Printf("n: %d, ways: %d\n", 0, climbStairs(0))
	fmt.Printf("n: %d, ways: %d\n", 1, climbStairs(1))
	fmt.Printf("n: %d, ways: %d\n", 2, climbStairs(2))
	fmt.Printf("n: %d, ways: %d\n", 3, climbStairs(3))
	fmt.Printf("n: %d, ways: %d\n", 4, climbStairs(4))
	fmt.Printf("n: %d, ways: %d\n", 5, climbStairs(5))
	fmt.Printf("n: %d, ways: %d\n", 6, climbStairs(6))
	fmt.Printf("n: %d, ways: %d\n", 45, climbStairs(45))
}
