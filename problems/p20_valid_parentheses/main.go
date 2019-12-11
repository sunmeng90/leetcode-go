package main

func main() {
	println(isValid(""))
	println(isValid("()"))
	println(isValid("()()"))
	println(isValid("(()[]){()}"))

}

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Complexity analysis

Time complexity : O(n) because we simply traverse the given string one character at a time and push and pop operations on a stack take O(1)O(1) time.
Space complexity : O(n) as we push all opening brackets onto the stack and in the worst case, we will end up pushing all the brackets onto the stack. e.g. ((((((((((.
*/
func isValid(s string) bool {
	stack := newStack()
	for i := 0; i < len(s); i++ {
		if stack.isEmpty() && isRight(s[i]) {
			return false
		} else {
			if isLeft(s[i]) {
				stack.push(s[i])
			} else if isRight(s[i]) && isPair(stack.peek(), s[i]) {
				stack.pop()
			} else {
				return false
			}
		}
	}
	return len(stack.arr) == 0
}

func isLeft(b byte) bool {
	return b == byte('(') || b == byte('{') || b == byte('[')
}

func isRight(b byte) bool {
	return b == byte(')') || b == byte('}') || b == byte(']')
}

func isPair(x, y byte) bool {
	return (x == byte('(') && y == byte(')')) || (x == byte('{') && y == byte('}')) ||
		(x == byte('[') && y == byte(']'))
}

type stack struct {
	arr []byte
}

func newStack() *stack {
	return &stack{
		arr: make([]byte, 0),
	}
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack) push(data byte) {
	s.arr = append(s.arr, data)
}

func (s *stack) pop() byte {
	if len(s.arr) == 0 {
		return 0
	}
	v := s.arr[len(s.arr)-1]
	if len(s.arr) == 1 {
		s.arr = s.arr[0:0]
	} else {
		s.arr = s.arr[:len(s.arr)-1]
	}
	return v
}

func (s *stack) peek() byte {
	if len(s.arr) == 0 {
		return 0
	}
	return s.arr[len(s.arr)-1]
}
