package main

import (
	fmt "fmt"
)

//go 的函数形式如同func 函数名 (参数列表)(返回值){函数体}
func test(a, b int) (sub, sum int) {
	sub = a - b
	sum = a + b
	return
}
func main() {
	//switch
	switch num := 3; num {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Printf("case 2")
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("This is defalut")
	}
	//循环
	for i := 1; i < 10; i++ {
		fmt.Println("loop:", i)
	}
	//key value遍历
	arr := [4]string{"hello", "this", "is", "key-value search"}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println(test(1, 2))
	//匿名函数
	var mul_func = func(a, b int) int {
		mul := a * b
		return mul
	}
	fmt.Println(mul_func(1, 2))
	//闭包 匿名函数使用了局部变量
}
