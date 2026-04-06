package main

func main() {
	print(BulkAtoi([]string{"8", "kood", "-13"}))
}

func BulkAtoi(arr []string) []int {
	result := []int{}
	
	for _,v :=range arr {
		result= append(result, StrToInt(v))
	}
	return result

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