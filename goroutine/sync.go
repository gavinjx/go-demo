//采用 sync.WaitGroup 的方式来阻塞 main 协程，等待其他协程的执行结束
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var tag = 0
var wg sync.WaitGroup

func main() {

	wg.Add(1)

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("main tag: %d\n", tag)
	//test goroutine
	go do()

	wg.Wait()

	fmt.Printf("main tag: %d\n", tag)

}

func do() {
	tag++
	fmt.Printf("do tag: %d\n", tag)
	wg.Done()
}