package main

import "container/list"

/*
Implement the following operations of a queue using stacks.

	push(x) -- Push element x to the back of queue.
	pop() -- Removes the element from in front of queue.
	peek() -- Get the front element.
	empty() -- Return whether the queue is empty.
Example:

	MyQueue queue = new MyQueue();

	queue.push(1);
	queue.push(2);
	queue.peek();  // returns 1
	queue.pop();   // returns 1
	queue.empty(); // returns false
Notes:

You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).
*/
// https://leetcode.com/problems/implement-queue-using-stacks/discuss/64284/Do-you-know-when-we-should-use-two-stacks-to-implement-a-queue
// https://stackoverflow.com/questions/2050120/why-use-two-stacks-to-make-a-queue/2050402#2050402
type MyQueue struct {
	inputStack  *list.List
	outputStack *list.List
}

/** Initialize your inputStack structure here. */
func Constructor() MyQueue {
	return MyQueue{inputStack: list.New(), outputStack: list.New()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inputStack.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Peek()
	return this.outputStack.Remove(this.outputStack.Back()).(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.outputStack.Len() == 0 { // reverse the stack to another stack the top will be the queue front
		for this.inputStack.Len() > 0 {
			inputTop := this.inputStack.Back()
			inputPop := this.inputStack.Remove(inputTop)
			this.outputStack.PushBack(inputPop)
		}
	}
	return this.outputStack.Back().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.inputStack.Len() == 0 && this.outputStack.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
