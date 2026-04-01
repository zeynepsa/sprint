package main
import "fmt"

var newValue rune
var newRune int
var newStep int

func main() {
	fmt.Printf("%c", ShiftBy('a', 4))
}

func ShiftBy(r rune, step int) rune {
	newStep = step % 26
	newRune = int(r) - 97
	newValue = rune(97 + newRune + newStep)
	return newValue
}