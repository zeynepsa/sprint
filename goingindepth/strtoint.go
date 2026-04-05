package main

func main() {
	print(StrToInt("10203"))

}

func StrToInt(s string) int {
	result := 0
	sign := 1
	for i, d := range s {
		if i == 0 {
			if d == '-' {
				sign = -1
				continue
			} else if d == '+' {
				continue
			}
		}
		if d < '0' || d > '9' {
			return 0
		}
		result = result*10 + int(d -'0')
	}
	return result*sign
}
