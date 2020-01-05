package main

/*
Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

Example:

Input: 38
Output: 2
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
             Since 2 has only one digit, return it.
Follow up:
Could you do it without any loop/recursion in O(1) runtime?
*/
/*
https://zhidao.baidu.com/question/30953893.html?qbl=relate_question_0&word=%CA%FD%D1%A7%B9%AB%CA%BD%20%C8%FD%BA%E1%CF%DF
3.同余符号
含义
两个整数a，b，若它们除以整数m所得的余数相等，则称a，b对于模m同余
记作a≡b(mod m)
读作a同余于b模m，或读作a与b关于模m同余。
比如26≡14(mod 12)。
定义
设m是大于1的正整数，a，b是整数，如果m|(a-b)，则称a与b关于模m同余，记作a≡b(mod m)，读作a同余于b模m。

proof:
https://blog.csdn.net/ray0354315/article/details/53991199
数根(digital root)公式的推导

*/

func addDigits(num int) int {
	return 1 + (num-1)%9
}

func main() {
	println(addDigits(38))
}
