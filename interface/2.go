package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Id   int
	Name string
}

func main() {
	fmt.Println("=============")
	user := User{
		Id:   1,
		Name: "Tom",
	}
	var ui, pi interface{} = user, &user

	pi.(*User).Id = 2
	// ui.(User).Id = 3 //ui 拷贝出来的是临时对象，不能对其进行赋值

	fmt.Println(ui)
	fmt.Println(pi)

	fmt.Println("##############")

	var a interface{} = nil
	var b interface{} = (*int)(nil)
	type iface struct {
		itab, data uintptr
	}
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	// tab = nil, data = nil
	// tab 包含 *int 类型信息, data = nil
	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())





}
