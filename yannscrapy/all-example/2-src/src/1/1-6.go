package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Getname() string { //p代表结构体本身的实列，类似python中的self,这里p可以写为self
	fmt.Println(p.Name)
	return p.Name
}

func main() {
	var person1 = new(Person)
	person1.Age = 34
	person1.Name = "darren"
	person1.Getname() // darren
}
