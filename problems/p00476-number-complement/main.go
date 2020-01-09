package main

/*
Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.

Note:
The given integer is guaranteed to fit within the range of a 32-bit signed integer.
You could assume no leading zero bit in the integer’s binary representation.
Example 1:
Input: 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
Example 2:
Input: 1
Output: 0
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
*/
func findComplement(num int) int {
	n := ^0         // all 1
	for n&num > 0 { // until the most significant bit of num (left most 1)
		n <<= 1
	} // the count of zero bits in n equals the number of bits from most significant bit of num to the end
	return ^n ^ num // revere 0 for n as 1 and reverse the bit for num
}

//0
//00000000000000000000000000000000
//^0
//11111111111111111111111111111111
//10
//00000000000000000000000000001010
//n<<=1 ...
//11111111111111111111111111110000
//^n
//00000000000000000000000000001111
//n ^ 10
//00000000000000000000000000000101
