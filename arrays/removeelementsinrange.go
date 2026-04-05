package main
import "fmt"
func main() {
	fmt.Print(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1))
}

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
		newArr :=[] float64{}
		if to<=from {
			from,to = to,from 
		}

		for i:=0 ; i<len(arr) ; i++ {
			if i<from || i>= to {
				newArr =append(newArr, arr[i])
		}
	}
		return newArr
}