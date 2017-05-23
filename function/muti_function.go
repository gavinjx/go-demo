package main

//函数多返回值
import "fmt"

func main() {
	x, y := double(2)
	fmt.Println(x, y)
}
func double(x int) (int, int) {
	return x, 2 * x
}
