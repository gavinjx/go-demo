package main

import "fmt"

func main() {
	getType := func(i interface{}) {

		switch t := i.(type) {
		case bool:
			fmt.Println("i am bool")
			break
		case int:
			fmt.Println("i am int")
			break
		case string:
			fmt.Println("i am string")
			break
		default:
			fmt.Println("i don`t know", t)
		}
	}
	getType(1)
	getType(true)
	getType("hello")
}
