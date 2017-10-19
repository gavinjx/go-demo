package main

import "fmt"

type I interface {
	Get() int
	Put(int)
}

type S struct {
	i int
}

type T struct {
	i int
}

func (p *S) Get() int {
	return p.i
}
func (p *S) Put(val int) {
	p.i = val
}

func GetData(i I) int {
	return i.Get()
	return 1
}

func main() {
	s := &S{}
	s.Put(12)
	fmt.Print(GetData(s))
}
