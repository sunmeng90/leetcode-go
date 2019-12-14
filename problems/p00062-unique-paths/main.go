package main

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:

	Input: m = 3, n = 2
	Output: 3
	Explanation:
	From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
	1. Right -> Right -> Down
	2. Right -> Down -> Right
	3. Down -> Right -> Right

Example 2:

	Input: m = 7, n = 3
	Output: 28

*/
func main() {
	//println(uniquePaths(51, 9)) //1916797311
	println(uniquePaths3(51, 9))
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	return uniquePaths(m-1, n) + uniquePaths(m, n-1)

}

// dynamic programming
func uniquePaths2(m int, n int) int {
	// path[i][j] 代表了，到达 (i,j) 格子的不同路径数目
	path := [100][100]int{}

	for i := 0; i < m; i++ {
		// 到达第 0 列的格子，只有一条路径
		path[i][0] = 1
	}

	for j := 0; j < n; j++ {
		// 到达第 0 行的格子，只有一条路径
		path[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 到达 (i,j) 格子的路径数目，等于
			// 到达 上方格子 和 左边格子 路径数之和
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	//for _, arr2 := range path {
	//	fmt.Printf("%v\n", arr2)
	//}
	return path[m-1][n-1]
}

func uniquePaths3(m int, n int) int {
	// path[i][j] 代表了，到达 (i,j) 格子的不同路径数目
	path := make([][]int, m, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n, n)
		path[i][0] = 1
	}
	for i := 0; i < n; i++ {
		path[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 到达 (i,j) 格子的路径数目，等于
			// 到达 上方格子 和 左边格子 路径数之和
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[m-1][n-1]
}
