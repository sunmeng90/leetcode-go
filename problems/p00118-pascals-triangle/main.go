package main

/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
	[
		 [1],
		[1,1],
	   [1,2,1],
	  [1,3,3,1],
	 [1,4,6,4,1]
	]
*/

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

func main() {
	result := generate(5)
	println(result)
}
