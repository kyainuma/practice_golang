package main

import "fmt"

func do(i interface{}) {
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	// ss := i.(string)
	// fmt.Println(ss)

	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't Know %T\n", v)
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)

	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}
