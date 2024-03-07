package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var wg1 sync.WaitGroup

func main() {
	wg.Add(2)
	// 多线程
	go foo()
	go bar()
	// wg.Wait()会阻塞，直到wg中counter为0
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
