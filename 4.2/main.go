package main

import "fmt"

var top int

var s [1000]int

func push(x int) {
	s[top] = x
	top = top + 1
}

func pop() int {
	top --
	return s[top]
}

func main() {
	var a, b int;
	top = 0

	arg := []interface{}{1, 2, "+", 3, 4, "-", "*"}

	for _, v := range arg {
		if num, ok := v.(int); ok {
			push(num)
		} else {
			op := v.(string)
			b = pop()
			a = pop()
			switch op {
			case "+":
				push(a + b)
			case "*":
				push(a * b)
			case "-":
				push(a - b)
			}
		}
	}

	//-3
	fmt.Println("rel ", pop())
}
