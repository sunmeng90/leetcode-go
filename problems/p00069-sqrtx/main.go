package main

/*
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

	Input: 4
	Output: 2

Example 2:

	Input: 8
	Output: 2
	Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/

/**
not working
*/
func mySqrt2(x int) int {
	if x == 1 {
		return 1
	}
	low, high, mid := 0.0, float64(x), float64(x)/2.0
	precision := 0.00001
	for high-low >= precision {
		if mid*mid > float64(x) {
			high = mid
		} else {
			low = mid
		}
		mid = (low + high) / 2
	}
	return int(mid)
}

func mySqrt(x int) int {
	i := 0
	j := (x / 2) + 1
	for i <= j {
		mid := (i + j) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			i = mid + 1
		}
		if mid*mid > x {
			j = mid - 1
		}
	}
	return j
}

func main() {
	println(mySqrt2(0))
	println(mySqrt2(1))
	println(mySqrt2(2))
	println(mySqrt2(3))
	println(mySqrt2(4))
	println(mySqrt2(5))
	println(mySqrt2(6))
	println(mySqrt2(7))
	println(mySqrt2(8))
	println(mySqrt2(9))
	println(mySqrt2(36))
	println(mySqrt(36))
}
