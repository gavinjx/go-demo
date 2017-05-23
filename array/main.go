package main

import "fmt"

const (
	name = "gavin"
)

func main() {
	var a [3]int
	a[0] = 1
	fmt.Println(a)
	fmt.Println(a[0], len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)
}
