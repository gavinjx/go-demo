package main

//闭包
import "fmt"

func seq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
func main() {
	step := seq()
	fmt.Println(step())
	fmt.Println(step())

	step = seq()
	fmt.Println(step())
}
