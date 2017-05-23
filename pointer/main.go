package main

import "fmt"

func main() {
	i := 1
	zeroval(i)
	fmt.Println(i)

	zeropointer(&i)
	fmt.Println(i)
}

func zeroval(i int) {
	i = 0
}
func zeropointer(iptr *int) {
	*iptr = 0
}
