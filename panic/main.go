package main

import "fmt"

func main() {
	fmt.Println("start")
	// defer是服从先进后出的执行规则，所以捕捉panic的defer放在首位才能最后被执行，起到catch的效果
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("catch err||err=%v", err)
		}
	}()

	defer func() {
		fmt.Println("Test1")
		panic("Here")
		fmt.Println("Test2")
	}()
}
