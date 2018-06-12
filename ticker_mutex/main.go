package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main() started")
	//var mutex = &sync.Mutex{}
	a := ""
	b := ""

	ticker1 := time.NewTicker(500 * time.Millisecond)
	go func(){
		for  _ = range ticker1.C {

			//mutex.Lock()
			a += "0"
			//mutex.Unlock()
			//fmt.Println("ticker 1: ", a, " ", t)
		}
	}()

	ticker2 := time.NewTicker(150 * time.Millisecond)
	go func(){
		for _ = range ticker2.C {

			//mutex.Lock()
			b += "1"
			//mutex.Unlock()
			//fmt.Println("ticker 2: ", a, " ", t)
		}
	}()

	ticker3 := time.NewTicker(1000 * time.Millisecond)
	go func(){
		for _ = range ticker3.C {
			fmt.Println(a)
		}
	}()

	time.Sleep(100000 * time.Millisecond)
	ticker1.Stop()
	ticker2.Stop()
	fmt.Println("main() finished")
}