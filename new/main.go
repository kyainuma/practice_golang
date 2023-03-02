package main

import "fmt"

func main() {
	// ポインタ以外を返すものはmake
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	// ポインタを返すものはnew
	var p *int = new(int)
	fmt.Printf("%T\n", p)

	// var p *int = new(int)
	// fmt.Println(*p)
	// *p++
	// fmt.Println(*p)

	// var p2 *int
	// fmt.Println(p2)
	// *p2++ メモリを確保していないためエラーになる
}
