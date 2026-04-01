package main

func main() {

	print(IsLeapYear(2000))
}

func IsLeapYear(year int) bool {
	if year % 4 == 0 && year % 400 == 0{
		return true
	}else {
		return false
	}
}