// 空接口
package main

import "fmt"

// Go 语言为了避免用户重复定义很多空接口，它自己内置了一个，这个空接口的名字特别奇怪，叫 interface{}
/*
空接口里面没有方法，所以它也不具有任何能力，其作用相当于 Java 的 Object 类型，可以容纳任意对象，
它是一个万能容器。比如一个字典的 key 是字符串，但是希望 value 可以容纳任意类型的对象，
类似于 Java 语言的 Map 类型，这时候就可以使用空接口类型 interface{}。
*/
func main() {
	// 连续两个大括号，是不是看起来很别扭
	// 代码中 user 字典变量的类型是 map[string]interface{}，
	// 从这个字典中直接读取得到的 value 类型是 interface{}，需要通过类型转换才能得到期望的变量。
	var user = map[string]interface{}{
		"age":     30,
		"address": "Beijing Tongzhou",
		"married": true,
	}
	fmt.Println(user)
	// 类型转换语法来了
	var age = user["age"].(int)
	var address = user["address"].(string)
	var married = user["married"].(bool)
	fmt.Println(age, address, married)

	user["price"] = 5.5
	var price = user["price"].(float32) // ？报错?
	fmt.Println("user: ", user, price)
}
