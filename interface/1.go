/*
【GoLang笔记】浅析Go语言Interface类型的语法行为及用法
https://studygolang.com/articles/2652
*/
package main

import (
	"fmt"
)

type IpAddr [4]byte
type AddrName string

func (p IpAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", p[0], p[1], p[2], p[3])
}

func (p AddrName) String() string {
	return fmt.Sprintf("key %v", string(p))
}

//空接口
func Print(v interface{}) {
	fmt.Printf("%T, %v\n", v, v)
}

func main() {
	ips := map[AddrName]IpAddr{
		"localhost": {127, 0, 0, 1},
		"dns":       {8, 8, 8, 8},
	}
	for name, ip := range ips {
		fmt.Printf("%v: %v\n", name, ip)
	}
	Print("Hello World!")
}
