//采用 sleep 的方式来阻塞 main 协程，等待其他协程的执行结束
package main

import (
	"fmt"
	"runtime"
	"time"
)

var tag = 0

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("main tag: %d\n", tag)
	//test goroutine
	go do()
	time.Sleep(200*100)
	fmt.Printf("main tag: %d\n", tag)

}

func do() {
	tag++
	fmt.Printf("do tag: %d\n", tag)
}