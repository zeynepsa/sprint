package main

func Doop (a int,op string, b int) int {
	switch op {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		default:
			return 0	
	}
}