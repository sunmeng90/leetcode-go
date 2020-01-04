package main

/*
Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/
// TODO: not easy to understand
// https://leetcode.com/problems/count-primes/discuss/57588/My-simple-Java-solution
func countPrimes(n int) int {
	notPrime := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if notPrime[i] == false {
			count++
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}
	return count
}
