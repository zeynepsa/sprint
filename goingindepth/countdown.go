package main

func main(){
	print(Countdown(7))
}

func Countdown(n int) string {
	str:=""
	for i :=n ;i>=0 ; i=i-2 {
		str += string(rune(i)+ '0') + ", "
	}
	return str + "0!"
}