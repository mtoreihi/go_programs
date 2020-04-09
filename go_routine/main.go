package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	fmt.Println("Hello World!")
	var count = int(13)
	var ptr = &count
	fmt.Println("pointer value=", *ptr)
	*ptr = 100
	fmt.Println(count)

	go f()
	time.Sleep(2 * time.Second)
	fmt.Println("main function")
	wg.Wait()
}

func f() {
	fmt.Println("f function started")
	fmt.Println("f function finished")
	wg.Done()
}
