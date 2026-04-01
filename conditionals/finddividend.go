package main

func main() {
	print(FindDividend(5, 17, 4))

}

func FindDividend(from, to, divisor int) int {
	for i := from; i < to; i++ {
		if i%divisor == 0 {
			return i
		}
	}
	return -1
}
