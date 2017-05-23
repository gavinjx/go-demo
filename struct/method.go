package main

import "fmt"

type animal struct {
	kind, name string
}

func (ani *animal) getName() string {
	return ani.name
}
func (ani *animal) getKind() string {
	return ani.kind
}
func main() {
	ani := animal{kind: "cat", name: "tom"}
	fmt.Println("kind:", ani.getKind())
	fmt.Println("name:", ani.getName())

	a := &ani

	a.kind = "dog"
	fmt.Println("kind:", ani.getKind())
	fmt.Println("name:", a.getName())

}
