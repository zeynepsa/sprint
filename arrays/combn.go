package main

import "fmt"

func main() {
	fmt.Print(CombN(1))
}

func CombN(n int) []string {
	result:=[]string{}
	//using recursive
	var combine func (start int, current string) 
	combine = func (start int, current string) {
		if len(current) == n {
			result =append(result, current)
			return 
		}
		for d:=start ; d<=9 ; d++ {
			combine(d+1,current + fmt.Sprint(d))
		}
	} 
	combine(0,"")
	return result
}