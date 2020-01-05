package main

/*
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of
your product fails the quality check. Since each version is developed based on the previous version, all the versions
after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following
ones to be bad.

You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find
the first bad version. You should minimize the number of calls to the API.

Example:

Given n = 5, and version = 4 is the first bad version.

call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true

Then 4 is the first bad version.
*/

func firstBadVersion(n int) int {
	return binarySearch(1, n)
}

func binarySearch(low, high int) int {
	mid := (low + high) / 2
	if isBadVersion(mid) {
		if mid == 1 || !isBadVersion(mid-1) {
			return mid
		}
		return binarySearch(low, mid-1)
	} else {
		return binarySearch(mid+1, high)
	}
}

func firstBadVersionNonRecursive(n int) int {
	low, high := 1, n
	for low <= high {
		mid := low + (high-low)/2 // avoid overflow when low and high are int.MAX_VALUE
		if isBadVersion(mid) {
			// if current version is the first bad version, then it will be compensate by low = mid + 1
			high = mid - 1 // if current is bad then go to previous
		} else {
			low = mid + 1 // if current is not bad version, then go to next
		}
	}
	return low
}

func firstBadVersionNonRecursive2(n int) int {
	low, high := 1, n
	for low < high {
		mid := low + (high-low)/2 // avoid overflow when low and high are int.MAX_VALUE
		if isBadVersion(mid) {
			high = mid // keep the bad version index, wait for low come to the same position and exit the loop
		} else {
			low = mid + 1
		}
	}
	return low
}

// pre defined api
func isBadVersion(n int) bool {
	return false
}
