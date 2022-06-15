package main

import (
	"fmt"
	"strings"
	"time"
)

//首字母大写表示可以在其他包用，小写表示仅可以在本包内使用
func div(a, b int) int {
	if b == 0 {
		panic("除数不能为0")
	}
	return a / b
}
func Stringtest() {
	str := "汉字"
	arr1 := []byte(str)
	//正常对于UTF-8来说，一个汉子占据3个字符
	fmt.Println(len(arr1))
	//rune数据结构类型用于保存汉字
	arr := []rune(str)
	fmt.Println(len(arr))
	var str1 string = "abcdefghijklmnopqrstuvwxzy"
	fmt.Println(strings.Index("a", str1))
	fmt.Println(strings.Contains(str1, "abc"))
	var str2 string = "bcdefghijklmnopqrstuvwxzy"
	fmt.Println(strings.Compare(str1, str2))
}
func masterbate() {
	for i := 1; i < 10; i++ {
		fmt.Println("I am masterbating")
		time.Sleep(time.Millisecond)
	}
	fmt.Println("I have cummed")
}
func dance() {
	for i := 1; i < 10; i++ {
		fmt.Println("I am dancing")
		time.Sleep(time.Millisecond)
	}
	fmt.Println("I have been shot")
}

/*
func main() {
	fmt.Println(div(1, 9))
	Stringtest()
	go masterbate()
	go dance()
	for {

	}
}
*/
