package main

import "strings"

/*
You are given a license key represented as a string S which consists only alphanumeric character and dashes. The string
is separated into N+1 groups by N dashes.

Given a number K, we would want to reformat the strings such that each group contains exactly K characters, except for
the first group which could be shorter than K, but still must contain at least one character. Furthermore, there must be
a dash inserted between two groups and all lowercase letters should be converted to uppercase.

Given a non-empty string S and a number K, format the string according to the rules described above.

Example 1:
Input: S = "5F3Z-2e-9-w", K = 4

Output: "5F3Z-2E9W"

Explanation: The string S has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.
Example 2:
Input: S = "2-5g-3-J", K = 2

Output: "2-5G-3J"

Explanation: The string S has been split into three parts, each part has 2 characters except the first part as it could
be shorter as mentioned above.

*/

func licenseKeyFormatting(S string, K int) string {
	count := 0
	result := ""
	for i := len(S) - 1; i >= 0; i-- {
		if '-' != S[i] {
			if count == K {
				result = "-" + result
				count = 0
			}
			result = string(S[i]) + result
			count++
		}
	}
	return strings.ToUpper(result)
}

// lower time and space cost
func licenseKeyFormatting2(S string, K int) string {
	count := 0
	result := make([]byte, 0)
	for i := len(S) - 1; i >= 0; i-- {
		if '-' != S[i] {
			if count == K {
				result = append(result, '-')
				count = 0
			}
			result = append(result, S[i])
			count++
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return strings.ToUpper(string(result))
}

// TODO: replace all '-' with "" and then insert '-' from end with step i-k

func main() {
	println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	println(licenseKeyFormatting("2-4A0r7-4k", 4))

	println(licenseKeyFormatting("--a-a-a-a--", 2))
	println(licenseKeyFormatting("0123456789", 1))
}
