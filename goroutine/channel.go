//采用 channel 的方式来阻塞 main 协程，等待其他协程的执行结束
package main

import (
	"fmt"
	"runtime"
	"time"
)

var tag = 0
var ch = make(chan bool)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("main tag: %d\n", tag)
	//test goroutine
	go do()

	var chVal = <- ch

	if(chVal){
		//其他协程的 ch
		fmt.Printf("chan\n")
		time.Sleep(200*100)
	}

	fmt.Printf("main tag: %d\n", tag)

}

func do() {
	ch <- true
	tag++
	fmt.Printf("do tag: %d\n", tag)
}