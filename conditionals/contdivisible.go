package main

func main() {
	print(CountDivisible(5, 17, 2, 3))
}

func CountDivisible(from, to, step, divisor int) int {
	count := 0
	if step <= 0 || divisor == 0 {
		return 0
	} else {
		for i := from; i < to; i = i + step {
			if i%divisor == 0 {
				count++
			}
		}
		return count
	}
}