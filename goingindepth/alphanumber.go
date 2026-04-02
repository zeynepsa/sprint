package main

func main() {
	print(AlphaNumber(0))
}

func AlphaNumber(n int) string {
		digits:= []int{}
		alphanumber:=""
			if n== 0 {
				return "a"
			}

			if n<0{
			alphanumber = "-"
			n=n*-1
		}

		for n>0 {
			digits=append([]int{n%10}, digits...)
			n/=10
		}

		for i:=0 ; i<len(digits); i++{
			alphanumber += string(rune(digits[i])+97)
		}

		return alphanumber

}