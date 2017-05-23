package main

import "fmt"

//可变函数
func main() {
	sum(1, 2, 3)
}

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("sum:", total)
}
