package main

//struct
import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"gavin", 25})
	fmt.Println(person{name: "gavin", age: 25})
	fmt.Println(person{name: "gavin"})
	fmt.Println(person{age: 25})
	fmt.Println(&person{name: "gavin", age: 25})
	s := person{name: "gavin", age: 24}
	fmt.Println(s)

	sp := &s
	fmt.Println("sp.name:", sp.name)
	sp.age = 21
	fmt.Println(s.age)
}
