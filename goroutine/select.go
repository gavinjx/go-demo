//采用 select 的方式来阻塞 main 协程，等待其他协程的执行结束
package main

import (
	"fmt"
	"runtime"
)

var tag = 0

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("main tag: %d\n", tag)

	//test goroutine
	go do()
	select {}

	fmt.Printf("main tag: %d\n", tag)

}

func do() {
	tag++
	fmt.Printf("do tag: %d\n", tag)
}
