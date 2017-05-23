package main

import "fmt"

func main(){
	m := make(map[string]string)
	fmt.Println(m)
	m["name"] = "gavin"
	m["city"] = "beijing"
	v := m["name"]
	fmt.Println(v)

	delete(m, "city")

	fmt.Println(m)

	k, val := m["name"]
	fmt.Println("k:", k, " v:", val)

	k, val = m["age"]
	fmt.Println("k:", k, " v:", val)


	newMap := map[string]int{
		"age": 18,
		"cityId":1,
	}
	fmt.Println(newMap)

}
