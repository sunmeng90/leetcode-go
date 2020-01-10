package main

/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I'll tell you whether the number is higher or lower.

You call a pre-defined API guess(int num) which returns 3 possible results (-1, 1, or 0):

-1 : My number is lower
 1 : My number is higher
 0 : Congrats! You got it!
Example :

Input: n = 10, pick = 6
Output: 6
*/

// mock provided api
func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	return doGuessIn(1, n)
}

func doGuessIn(low, high int) int {
	mid := (low + high) / 2
	guessIndicator := guess(mid)
	if guessIndicator == 0 {
		return mid
	} else if guessIndicator == -1 {
		return doGuessIn(low, mid-1)
	} else {
		return doGuessIn(mid+1, high)
	}
}
