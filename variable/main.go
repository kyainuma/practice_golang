package main

import "fmt"

// var は関数外でも宣言可能
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func variables() {
	// ショートは関数内で宣言可能
	xi := 1
	xf64 := 1.2
	// 明示的に型指定したい場合はvar
	var xf32 float32 = 3.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)

}
func main() {
	fmt.Println(i, f64, s, t, f)
	variables()
}
