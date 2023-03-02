package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println("world")
	// fmt.Println("hello")

	// fmt.Println("run")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")

	file, _ := os.Open("./test.txt")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
