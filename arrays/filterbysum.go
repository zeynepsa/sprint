package main

import "fmt"

func main() {
	fmt.Print(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 4))
}

func FilterBySum(arr [][]int, limit int) [][]int {
	
	result := [][]int{}

	for i, _ := range arr {
		sum := 0
		for _, y := range arr[i] {
			sum += y
		}
		if sum >= limit {
				result = append(result, arr[i])
			}
		
	}
	return result
}
