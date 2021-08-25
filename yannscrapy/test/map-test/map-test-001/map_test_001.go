package main

import "fmt"

type Student struct {
	Name string
	Age  int
	Addr string
}

func main() {
	list := make([]*Student, 0)
	m := make(map[string]*Student, 0)

	s1 := &Student{
		Name: "lilei",
		Age:  18,
		Addr: "beijing",
	}

	s2 := &Student{
		Name: "alien",
		Age:  19,
		Addr: "shanghai",
	}

	list = append(list, s1)
	list = append(list, s2)

	m["lilei"] = s1
	m["alien"] = s2

	s3 := m["lilei"]
	s3.Age = 20
	s3.Addr = "guangzhou"

	fmt.Printf("list:\n")
	for _, v := range list {
		fmt.Printf("v:%v\n", v)
	}

	fmt.Printf("m:\n")
	for _, v := range m {
		fmt.Printf("v:%v\n", v)
	}
}
