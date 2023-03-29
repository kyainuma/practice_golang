package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// このメソッドは、型TがインターフェースIを実装していることを意味します
// しかし、そうであることを明示的に宣言する必要はありません。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}