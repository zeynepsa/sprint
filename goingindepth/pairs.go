package main
import "fmt"

func main () {
	fmt.Printf(Pairs())
}

func Pairs() string {
	str :=""
	for i := 0; i <100; i++ {
		for y:=i+1; y<100;y++ {
			str = str + fmt.Sprintf("%02d %02d, ", i, y)
		}
}
		return str[:len(str)-2]
}