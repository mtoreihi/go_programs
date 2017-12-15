package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("FolderList.txt")
	if (err != nil) {
		panic(err)
	}
	defer file.Close()
	
	var line string
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		line = scanner.Text()
		fmt.Println(line)
		os.Mkdir(line, os.ModePerm)
		os.Mkdir(line + "\\Logs", os.ModePerm)
		os.Mkdir(line + "\\Requests", os.ModePerm)
		//os.Exit(1)
	}



}