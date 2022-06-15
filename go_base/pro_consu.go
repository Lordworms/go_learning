package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var lock = sync.Mutex{}
var sce []int = make([]int, 10)

func producer() {
	lock.Lock()
	for i := 1; i < 10; i++ {
		num := rand.Intn(100)
		sce[i] = num
		fmt.Println("producer has produced:", num)
	}
	lock.Unlock()
}
func consumer() {
	lock.Lock()
	for i := 1; i < 10; i++ {
		num := sce[i]
		fmt.Println("consumer has consumed", num)
	}
	lock.Unlock()
}

/*
func main() {
	go producer()
	go consumer()
	for {

	}
}
*/
