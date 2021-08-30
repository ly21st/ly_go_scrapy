package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int64
	wight int64
	high  int64
	score int64
}

func main() {
	var stu1 = new(Student)
	fmt.Printf("地址分布:")
	fmt.Printf("%p\n", &stu1.Name)
	fmt.Printf("%p\n", &stu1.Age)
	fmt.Printf("%p\n", &stu1.wight)
	fmt.Printf("%p\n", &stu1.high)
	fmt.Printf("%p\n", &stu1.score)
	typ := reflect.TypeOf(Student{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i) // 反射出filed
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
}
