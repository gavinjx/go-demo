package main
import (
	"fmt"
)
func main(){
	var a,b int32
	a = 7
	b = 0
	//判断奇偶
	fmt.Println(a&1)
	fmt.Println(b&1)
	fmt.Println("=================")

	var aa,bb int32
	aa = -7
	bb = -8
	//判断奇偶
	fmt.Println(aa&1)
	fmt.Println(bb&1)
	fmt.Println("=================")

	//除以2
	var d uint32
	d = 22
	fmt.Println(d >> 1)
	fmt.Println("=================")

	//乘以2
	var dd uint32
	dd = 22
	fmt.Println(dd << 1)
	fmt.Println("=================")

	//异或取整数最大值
	const MaxUint = ^uint(0)
	fmt.Println(MaxUint)
	fmt.Println("=================")

	//有符号右移：算术右移
	var c int8
	c = -5
	fmt.Printf("%b\n", c)
	fmt.Printf("%b\n", c>>1)
	fmt.Println(c>>1)
	fmt.Println("=================")

	//无符号右移
	var e uint8
	e = 5
	fmt.Printf("%b\n", e)
	fmt.Println(e>>1)
	fmt.Println("=================")

}
