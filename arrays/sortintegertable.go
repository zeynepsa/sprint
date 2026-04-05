package main

import "fmt"

func main() {
	fmt.Print(SortIntegerTable([]int{2, 0, 5, 4, 1, 3}))
}

func SortIntegerTable(table []int) []int {

	for i := range table[:len(table)-1] {
			for j:=i+1 ;j<len(table) ; j++{
				if table[j]<table[i] {
					table[j], table[i] = table[i], table[j]
				}
			}
	}
	return table
}
