package main

import (
	"unsafe"
	"fmt"
)

func main(){
	d := struct {
		s string
		x int
	}{"abc",100}

	up := unsafe.Pointer(&d)
	p := uintptr(up)
	// 对struct d 的指针偏移到 d.x
	p+=unsafe.Offsetof(d.x)

	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	// 对 d.x 进行赋值
	*px = 200
	fmt.Printf("%#v\n", d)
}