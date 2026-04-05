package main


func main() {
	print(SimpleStrToInt("10203"))

}

func SimpleStrToInt(s string) int {
	digits := []rune(s)
	var d1 int
	sum := 0

	//checking if invalid input is given
	for _, d := range s {
		if d < '0' || d > '9' {
			return 0
		}
	}
	//to find which digit the number actually starts
	for i, d := range s {
		if d != '0' {
			d1 = i
			break
		}
	}

	for i := d1; i < len(digits); i++ {
		number := 1
		for j := 0; j < len(digits)-1-i; j++ {
			number *= 10
		}
		sum += int(digits[i]-'0') * number
	}
	return sum
}
