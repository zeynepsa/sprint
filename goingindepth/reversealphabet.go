package main

func main () {
	print(ReverseAlphabet(5))
}
func ReverseAlphabet(step int) string {
	str := ""
	if step <= 0 {
		step = 1
	} 
		for i := 'z' ; i >= 'a'; i -= rune(step){
			str = str + string(i)
		}
		return str
}