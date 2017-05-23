package main

//递归
import "fmt"

func recursion(n int) int {
	fmt.Println(n)
	if n > 1 {
		return recursion(n - 1)
	} else {
		return 0
	}

}
func main() {
	recursion(6)
}
