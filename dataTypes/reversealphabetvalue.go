package main
import "fmt"

func main() {
	fmt.Printf("%c", ReverseAlphabetValue('b'))
}

func ReverseAlphabetValue(ch rune) rune {
		return 'z'- (ch- 'a')
}