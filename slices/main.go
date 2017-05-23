package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s, len(s), len(s[1]), s[1])

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	l = s[2:]
	fmt.Println(l)

	t := []string{ "1", "2", "3"}
	fmt.Println(t)

	for k, v := range t{
		fmt.Println("range:", k, v)
	}

	 o := make([][]int, 3)
	for i:=0; i<3; i++{
		o[i] = make([]int, i+1)
		for j:=0; j<len(o[i]); j++{
			o[i][j] = j+1
		}
	}
	fmt.Println(o)

}
