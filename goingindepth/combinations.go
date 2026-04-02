package main

import "fmt"

func main() {
	fmt.Printf(Combinations())
}

func Combinations() string {
	list := ""
	for num1 := 0; num1 < 10; num1++ {
		for num2 := num1 + 1; num2 < 10; num2++ {
			for num3 := num2 + 1; num3 < 10; num3++ {
				list += fmt.Sprintf("%v%v%v, ", num1, num2, num3)
			}
		}

	}
	return list
}
