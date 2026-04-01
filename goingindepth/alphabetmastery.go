package main

func main() {
	print(AlphabetMastery(6))
}

func AlphabetMastery(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str = str + string(rune(i+'a'))
	}
	return str
}