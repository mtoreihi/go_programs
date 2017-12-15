package main

import "fmt"
import "time"


var myarray [5]int

func main() {
	//var array [5]int
	myarray := [5]int{1,2,3,4,5}
	fmt.Println("Hello from Mehran", myarray[3], myarray[2], " ending...")
	testArray(&myarray)
	time.Sleep(10000 * time.Millisecond)
}

func testArray(array *[5]int) {
	fmt.Println(array[2])
}
