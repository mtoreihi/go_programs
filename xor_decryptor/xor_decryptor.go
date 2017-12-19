package main

import "os"
import "bufio"

func main() {
	f, err := os.Open("myfile.enc")
	if err != nil {
		panic("error opening file")
	}
	defer f.Close()
	o, err := os.Create("myfile.dec")
	if err != nil {
		panic("error opening file")
	}
	defer o.Close()

	var key byte = 0x13
	b1 := make([]byte, 1)
	var b2 byte
	w := bufio.NewWriter(o)

	for true {
		n, _ := f.Read(b1)
		if n != 1 {
			break
		}
		b2 = b1[0] ^ key
		w.WriteByte(b2)
	}
	w.Flush()
}
