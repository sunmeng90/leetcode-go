package main

import "fmt"

/*
Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.

Note:

All letters in hexadecimal (a-f) must be in lowercase.
The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero
character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
The given number is guaranteed to fit within the range of a 32-bit signed integer.
You must not use any method provided by the library which converts/formats the number to hex directly.
Example 1:

Input:
26

Output:
"1a"
Example 2:

Input:
-1

Output:
"ffffffff"
*/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hexChars := make([]byte, 0)
	for i := 0; i < 8 && num != 0; i++ {
		hc := num & 15
		if hc < 10 {
			hexChars = append(hexChars, byte('0'+hc))
		} else {
			hexChars = append(hexChars, byte('a'+hc-10))
		}
		num >>= 4
	}
	for i, j := 0, len(hexChars)-1; i < j; i, j = i+1, j-1 {
		hexChars[i], hexChars[j] = hexChars[j], hexChars[i]
	}
	return string(hexChars)
}

func main() {
	//println(toHex(-1))
	//println(toHex(0))
	//println(toHex(9))
	//println(toHex(10))
	//println(toHex(11))
	//println(toHex(12))
	//println(toHex(13))
	//println(toHex(14))
	//println(toHex(15))
	//println(toHex(16))
	//println(toHex(17))
	println(toHex(111111))
	println(toHex(0))
	//println(toHex(01111111111111111111111111111111))
	println("---------------------")
	fmt.Printf("%x\n", 17)
	fmt.Printf("%x\n", -1)
	fmt.Printf("%x\n", 111111)
}
