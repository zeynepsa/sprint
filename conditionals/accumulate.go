package main

func Accumulate(n int) int {
	sum := 0

	if n < 0 {
		return 0
	} else {
		for i := 0; i < n; i++ {
			sum = sum + i
		}
		return sum
	}
}