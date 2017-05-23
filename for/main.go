package main

import "fmt"

//for 循环使用的例子

var x = 0

func main() {
	//无限制循环下去，等于： for true { ... }
	//for {
	//	x++
	//	fmt.Printf("x: %d\n", x)
	//}

	//for x<10 {
	//	x++
	//	fmt.Printf("x: %d\n", x)
	//}

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d\n", i)
	}

}
