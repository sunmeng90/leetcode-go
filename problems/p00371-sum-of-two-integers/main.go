package main

/*
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:

Input: a = 1, b = 2
Output: 3
Example 2:

Input: a = -2, b = 3
Output: 1
*/
// https://www.cs.cornell.edu/~tomf/notes/cps104/twoscomp.html
// https://leetcode.com/problems/sum-of-two-integers/discuss/84290/Java-simple-easy-understand-solution-with-explanation
// https://blog.csdn.net/u014496330/article/details/53911427
// 二进制运算也是一位一位地相加，如果两个数当前位都为1，则左边一位加上1
// 两个数的异或值，也就是不需要进位的那些位的值
// 两个数的取与的值向左移一位，也就是进位后的值
// 一直重复这个操作，直到进位制为0，也就是没有进位啦，那么上面的数就是结果啦
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for b != 0 {
		carry := a & b // 加法中进位: 0,0 -> 0, 1,1->0(carry 1), 1,0->1, 0,1->1
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {

}
