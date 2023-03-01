package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	// アスキー文字で出力される
	fmt.Println("Hello World"[0])
	// stringで変換する
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	// これはできない
	// s[0] = "X"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println(`Test
				Test
Test`)
}
