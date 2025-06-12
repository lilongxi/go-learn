package main

import "fmt"

func main() {
	const name = "zhangsan"
	fmt.Println(name)

	const PI float32 = 3.1415926
	const P2 = 3.14
	const MY_NAME = "zhangsan"
	fmt.Println(PI, P2, MY_NAME)

	const (
		OK        = 200
		NOT_FOUND = 404
	)
	fmt.Println(OK, NOT_FOUND)

	const CODE1, CODE2 = 100, 200

	fmt.Println(OK, CODE1, MY_NAME)

	// 类型推断 未赋值的会延用前边的值
	const (
		x int = 16
		y
		y2
		s = ""
		z
	)

	fmt.Println(y2)
}
