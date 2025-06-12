package main

import "fmt"

// iota 特殊常量 可以被编译器修改的常量
// 类似枚举
func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)

	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)

	const (
		ERR1 = iota
		ERR2
		ERR3 = "lilongxi"
		ERR4
		ERR5 = iota
	)

	fmt.Println(ERR1, ERR2, ERR3, ERR4, ERR5)
}
