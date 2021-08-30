// 1.4 struct匿名成员（字段、属性)

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Student struct {
	score  string
	Age    int
	Person // 匿名内嵌结构体
}

func main() {
	var stu = new(Student)
	stu.Age = 22                         //优先选择Student中的Age
	fmt.Println(stu.Person.Age, stu.Age) // 0,22
}
