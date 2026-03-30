package dataTypes

func main() {
	print(mean(3.5, 1.5, 5.0))
}

func mean(a, b, c float32) float32 {
	return (a + b + c) / 3
}
