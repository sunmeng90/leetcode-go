package main

/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.

*/
func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1 // assign short string to num1
	}
	result := make([]byte, len(num2))
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; j >= 0; i, j = i-1, j-1 {
		digitSum := num2[j] + uint8(carry)
		if i >= 0 {
			digitSum = digitSum + num1[i] - '0'
		}
		if digitSum <= '9' {
			carry = 0
			result[j] = digitSum
		} else {
			carry = 1
			result[j] = digitSum - 10
		}
	}
	if carry == 1 {
		return "1" + string(result)
	}
	return string(result)
}

func main() {
	println(addStrings("0", "0"))
	println(addStrings("0", "1"))
	println(addStrings("1", "1"))
	println(addStrings("11", "1"))
	println(addStrings("91", "9"))
}
