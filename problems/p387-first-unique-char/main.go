package main

/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

	s = "leetcode"
	return 0.

	s = "loveleetcode",
	return 2.
Note: You may assume the string contain only lowercase letters.

*/

func firstUniqChar(s string) int {
	m := make(map[byte]int, 26) // frequency
	n := make(map[byte]int, 26) // index of first occurrence
	for i := 0; i < len(s); i++ {
		_, e := m[s[i]]
		if e {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
			n[s[i]] = i
		}
	}
	i := len(s)
	for k, v := range m {
		if v == 1 && n[k] < i {
			i = n[k]
		}
	}
	if i == len(s) {
		return -1
	}
	return i
}

func firstUniqChar2(s string) int {
	if len(s) == 0 {
		return -1
	}
	arr := make([]int, 26)
	for _, v := range s {
		arr[byte(v)-'a']++
	}

	for i, v := range s {
		if arr[byte(v)-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar3(s string) int {
	magic := [26]int{}
	for i := range magic {
		magic[i] = -1
	}
	for i, v := range s {
		if magic[v-'a'] == -1 {
			magic[v-'a'] = i
		} else {
			magic[v-'a'] = -2
		}
	}
	ret := len(s) + 1
	for _, v := range magic {
		if v >= 0 {
			if v < ret {
				ret = v
			}
		}
	}
	if ret == len(s)+1 {
		ret = -1
	}
	return ret
}
