package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("hoge")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging")
	log.Printf("%T %v", "Printf", "Printf")
	log.Fatalf("%T %v", "Fatalf", "Fatalf")
	log.Fatalln("error")

	fmt.Println("OK?")
}
