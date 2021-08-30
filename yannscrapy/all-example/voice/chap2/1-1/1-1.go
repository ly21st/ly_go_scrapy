package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	name  string
	age   int32
	Class string
}

func main() {
	// 1 值形式
	var stu1 Student
	stu1.age = 34
	stu1.name = "darren"
	stu1.Class = "class1"
	fmt.Println(stu1.name) //darren

	// 2 new 函数创建
	var stu2 *Student = new(Student)
	stu2.name = "king"
	stu2.age = 33
	fmt.Println(stu2.name, (*stu2).name) //king

	// &形式创建
	var stu3 *Student = &Student{
		name:  "rose",
		age:   18,
		Class: "class3", // 如果分行的时候每行都要,
	}
	fmt.Println(stu3.name, (*stu3).name) //rose  rose

	// 值 初始化
	var stu4 Student = Student{ // KV 形式初始化值
		name:  "柚子老师",
		age:   18,
		Class: "Go", // 注意这里的逗号不能少
	}
	fmt.Println("stu4:", stu4)

	// 值顺序初始化
	var stu5 Student = Student{ // 顺序形式 形式初始化值
		"1",
		18,
		"音视频", // 注意这里的逗号不能少
	}
	fmt.Println("stu5:", stu5)

	// nil结构体
	var stu6 *Student = nil
	fmt.Println("stu6:", stu6)

	// 结构体大小
	fmt.Println("unsafe.Sizeof(stu5):", unsafe.Sizeof(stu5))
	fmt.Println("unsafe.Sizeof(stu6):", unsafe.Sizeof(stu6))

	// fmt.Println("unsafe.Sizeof(string):", unsafe.Sizeof(string))
	// fmt.Println("unsafe.Sizeof(int):", unsafe.Sizeof(int))
}
