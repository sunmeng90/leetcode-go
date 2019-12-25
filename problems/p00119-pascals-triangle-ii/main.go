package main

/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

	Input: 3
	Output: [1,3,3,1]
Follow up:

Could you optimize your algorithm to use only O(k) extra space?
*/
func getRow(rowIndex int) []int {
	table := generate(rowIndex + 1)
	return table[rowIndex]
}

func generate(numRows int) [][]int {
	table := make([][]int, numRows)
	populateRow(numRows, &table)
	return table
}

func populateRow(numRows int, table *[][]int) {
	row := make([]int, numRows)
	if numRows == 0 {
		return
	}
	if numRows == 1 {
		row[0] = 1
	} else {
		populateRow(numRows-1, table)
		preRow := (*table)[numRows-2]
		row[0], row[numRows-1] = 1, 1
		for i := 1; i < numRows-1; i++ {
			row[i] = preRow[i-1] + preRow[i]
		}
	}
	(*table)[numRows-1] = row
}
