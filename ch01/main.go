package main

import "fmt"

// var age = 100
var (
	g_age  = 100
	g_name = "zhangsan"
	g_ok   bool
)

/**
* 1. 变量需要先定义后使用
* 2. 变量必须有类型 并且必须使用
* 3. 类型定下来后不能改变
 */
func main() {
	fmt.Println("Hello, Go!")
	var name string
	fmt.Println(name, name == "")
	// var age int = 18
	// var age = 18
	age := 1
	fmt.Println(age)
	fmt.Println(g_age, g_name, g_ok)

	var user1, user2, user3 = "bobby1", 1, true
	fmt.Println(user1, user2, user3)
}
