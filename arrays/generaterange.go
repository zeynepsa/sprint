package main

func main() {
	print(GenerateRange(-1, 4))
}

func GenerateRange(min, max int) []int {
	ranges :=[]int{}
		if max>min {
		for i:=min; i<max;i++ {
			ranges= append(ranges, i)
		}
		return ranges
	}else {
		return ranges
	}
}