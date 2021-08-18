package main

import (
	"fmt"
)

func change(a int) {
	a = 50
}
func main() {
	a := 55
	b := a
	change(b)
	fmt.Println(a)
}
