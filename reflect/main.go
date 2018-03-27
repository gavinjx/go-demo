package main

// https://studygolang.com/articles/4988
import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `iamtag`
	Id   int    `type:"userid" range:100`
}

func say(say string) {
	fmt.Printf("=====\n%s\n", say)
}

func main() {
	// 接口读取field tag
	var u User
	t := reflect.TypeOf(u)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%+v\n", t.NumField())
	fmt.Printf("tag: %v\n", t.Field(0).Tag)
	fmt.Printf("tag: %v\n", t.Field(1).Tag)
	fmt.Printf("tag: %v\n", t.Field(1).Tag.Get("type"))

	// 接口实现函数调用
	var funcMap = make(map[string]func(string))
	funcMap["say"] = say
	funcMap["say"]("hello")

}
