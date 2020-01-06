package main

/*
Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

Example 1:

Input: 16
Output: true
Example 2:

Input: 5
Output: false
Follow up: Could you solve it without loops/recursion?
*/
func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	for num%4 == 0 {
		num /= 4
	}
	return num == 1
}

/*
Eg. 16 -> 10000
10000 & 0101 0101 0101 0101 0101 0101 0101 0101 ->
        0000 0000 0000 0000 0000 0000 0001 0000, true, not zero
1000 & 0111 -> 0000, true, is zero
return true

Eg. 32 -> 100000
10 0000 & 0101 0101 0101 0101 0101 0101 0101 0101 ->
          0000 0000 0000 0000 0000 0000 0000 0000, false, this should be non zero in order to be power of four.
10 0000 && (01 1111) -> 0, true, is zero
return false
*/
func isPowerOfFour2(num int) bool {
	//0x55555555 is to get rid of those power of 2 but not power of 4
	//so that the single 1 bit always appears at the odd position
	return num > 0 && num&(num-1) == 0 && (num&0x55555555) != 0
}

func isPowerOfFour3(num int) bool {
	//a -> 1010
	return num > 0 && num&(num-1) == 0 && (num&0xaaaaaaaa) == 0
}

/*
(4^n - 1) % 3 == 0
another proof:
(1) 4^n - 1 = (2^n + 1) * (2^n - 1)
(2) among any 3 consecutive numbers, there must be one that is a multiple of 3
among (2^n-1), (2^n), (2^n+1), one of them must be a multiple of 3, and (2^n) cannot be the one, therefore either
(2^n-1) or (2^n+1) must be a multiple of 3, and 4^n-1 must be a multiple of 3 as well.
*/
func isPowerOfFour4(num int) bool {
	return num > 0 && num&(num-1) == 0 && (num-1)%3 == 0
}
