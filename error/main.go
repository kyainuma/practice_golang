package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[0]
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(count, string(data))
}
