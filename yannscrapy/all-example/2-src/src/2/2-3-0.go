// 2.3 interface使用

package main

import "fmt"

//定义接口
type Skills interface {
	Running()
	Getname() string
}

type Student struct {
	Name string
	Age  int
}

// 实现接口
func (p Student) Getname() string { //实现Getname方法
	fmt.Println(p.Name)
	return p.Name
}

func (p Student) Running() { // 实现 Running方法
	fmt.Printf("%s running", p.Name)
}

func main() {
	var skill Skills
	var stu1 Student
	stu1.Name = "darren"
	stu1.Age = 34
	skill = stu1
	skill.Running() //调用接口
}
