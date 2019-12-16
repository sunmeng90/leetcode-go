package main

import (
	"fmt"
	"github.com/sunmeng90/leetcode-go/util"
)

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

	Input: a = "11", b = "1"
	Output: "100"

Example 2:

	Input: a = "1010", b = "1011"
	Output: "10101"
*/
func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	idxRes := util.Max(la, lb) - 1
	res := make([]byte, idxRes+1)
	overflow := false
	zero, one := "0"[0], "1"[0]
	for i, j := la-1, lb-1; i >= 0 && j >= 0; {
		if overflow {
			if a[i] == one && b[j] == one {
				res[idxRes] = one
			} else if a[i] == zero && b[j] == zero {
				res[idxRes] = one
				overflow = false
			} else {
				res[idxRes] = zero
			}
		} else {
			if a[i] == one && b[j] == one {
				res[idxRes] = zero
				overflow = true
			} else if a[i] == zero && b[j] == zero {
				res[idxRes] = zero
			} else {
				res[idxRes] = one
			}
		}
		i--
		j--
		idxRes--
	}
	if la > lb {
		for idxRes >= 0 {
			if overflow {
				if a[idxRes] == one {
					res[idxRes] = zero
				} else {
					res[idxRes] = one
					overflow = false
				}
			} else {
				res[idxRes] = a[idxRes]
				overflow = false
			}
			idxRes--
		}
	} else if la < lb {
		for idxRes >= 0 {
			if overflow {
				if b[idxRes] == one {
					res[idxRes] = zero
				} else {
					res[idxRes] = one
					overflow = false
				}
			} else {
				res[idxRes] = b[idxRes]
				overflow = false
			}
			idxRes--
		}
	}

	if overflow {
		return "1" + string(res)
	}

	return string(res)
}

func main() {
	fmt.Printf("0 + 0 = %s\n", addBinary("0", "0"))
	fmt.Printf("0 + 1 = %s\n", addBinary("0", "1"))
	fmt.Printf("1 + 1 = %s\n", addBinary("1", "1"))
	fmt.Printf("10 + 10 = %s\n", addBinary("10", "10"))
	fmt.Printf("11 + 11 = %s\n", addBinary("11", "11"))
	fmt.Printf("111 + 11 = %s\n", addBinary("111", "11"))
	fmt.Printf("1011 + 11 = %s\n", addBinary("1011", "11"))
	fmt.Printf("1 + 111 = %s\n", addBinary("1", "111"))
}
