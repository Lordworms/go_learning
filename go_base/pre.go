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

//golang中的delay是一个栈的形式
func delay_deploy() {
	defer fmt.Println("3") //后
	fmt.Println("1")
	defer fmt.Println(("2")) //先
}

//单个包的调用流程：main包-->常量-->全局变量-->init函数-->main函数-->Exit
/*
func init() {
	fmt.Println("我是init")

}
*/
type student struct {
	name string
	id   int
}

func pre() {
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
	//略了

	//延迟
	delay_deploy()

	//支持数组直接比较
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 4}
	fmt.Println(arr1 == arr2)

	//可变长数组slice
	/*
		type slice struct{
		array unsafe.Pointer // 指向底层数组指针
		len int // 切片长度(保存了多少个元素)
		cap int // 切片容量(可以保存多少个元素)
		}
		创建方式：
		1.数组切片创建arr[st:end]
		2.make创建
		3.不指定长度的数组 var sce:=[]int{}
	*/
	//类型+长度+容量，如果没有指定容量，则容量=长度,对slice直接就是=,append后容量扩充一倍
	//切片可以再次产生新的slice，底层指向同一个数据结构，切片不支持== !=判断
	var sce = make([]int, 3, 5)
	fmt.Println(len(sce))
	fmt.Println(cap(sce))
	var sce_source = []int{1, 2, 3}
	copy(sce, sce_source)
	sce = append(sce, 0)
	sce = append(sce, 0)
	sce = append(sce, 0)
	sce = append(sce, 0)
	println(cap(sce))
	/*map是字典
	创建方式有三种：
		1.快速创建：
			dict  := map[string]string{"name":"lnj", "age":"33", "gender":"male"}
		2.make创建
			var dict = make(map[string]string, 3)
		3.make指定容量创建
			var dict = make(map[string]string)
		和slice一样，没有make的map是无法使用的，因为没有分配空间
	*/
	var dic map[int]string = map[int]string{1: "aa", 2: "bb", 3: "ccc"}
	fmt.Println(dic)
	delete(dic, 1)
	fmt.Println(dic)
	//struct
	a := student{"dick", 1}
	fmt.Println(a)
	var ta = struct {
		name string
		age  int
	}{"aaa", 2}
	fmt.Println(ta)
}
