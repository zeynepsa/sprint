package main
import "fmt"

func main() {
	fmt.Print(BalanceOut([]bool{false, true, true, true}))
}

func BalanceOut(arr []bool) []bool {
	trueCount :=0
	falseCount :=0 

	for _,v :=range arr {
		if v { //instead of v==true 
			trueCount++
		}else {
			falseCount++
		}
	}

	diff :=trueCount-falseCount

	for i := 0; i < diff; i++ {
        arr = append(arr, false)
    }
    for i := 0; i < -diff; i++ {
        arr = append(arr, true)
    }

    return arr

	

}
