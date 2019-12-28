package main

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
	push(x) -- Push element x onto stack.
	pop() -- Removes the element on top of the stack.
	top() -- Get the top element.
	getMin() -- Retrieve the minimum element in the stack.


Example:
	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin();   --> Returns -3.
	minStack.pop();
	minStack.top();      --> Returns 0.
	minStack.getMin();   --> Returns -2.
*/
type MinStack struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.data = this.data[0 : len(this.data)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	min := 0
	if len(this.data) > 0 {
		min = this.data[0]
		for _, v := range this.data {
			if v < min {
				min = v
			}
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	var minStack MinStack = Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin() //   --> Returns -3.
	minStack.Pop()
	minStack.Top()    //   --> Returns 0.
	minStack.GetMin() //  --> Returns -2.
}
