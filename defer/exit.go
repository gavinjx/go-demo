package main
//defer 的语句是在 func 结束后执行的，但是在调用 os.Exit 之后并不会再执行 defer 的语句
import (
	"fmt"
	"os"
)

func main(){
	defer print()

	fmt.Println("start ...")
	os.Exit(0)
}

func print()  {
	fmt.Println("exit ...")
}