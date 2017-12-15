package main

import "crypto/md5"
import "fmt"
import "io"

func main() {
	h := md5.New()
	io.WriteString(h, "Mehran Toreihi")
	fmt.Println("MD5 sum:")
	fmt.Printf("%x", h.Sum(nil))
}