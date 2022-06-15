package main

import "fmt"

func main() {
	var mych chan int
	mych = make(chan int, 3)
	fmt.Println("length of mych is ", len(mych), ". cap of mych is", cap(mych))
	mych <- 666
	fmt.Println("length of mych is ", len(mych), ". cap of mych is", cap(mych))
	num := <-mych
	fmt.Println("num is:", num)
	fmt.Println("length of mych is ", len(mych), ". cap of mych is", cap(mych))
	mych <- 111
	mych <- 222
	mych <- 333
	for num, ok := range mych {
		fmt.Println(num, ok)
	}

}
