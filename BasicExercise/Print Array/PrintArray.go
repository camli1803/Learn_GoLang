package main

import (
	"fmt"
)

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
func main() {
	a := [3][2]string{
		{"CL", "MC"},
		{"DT", "TP"},
		{"PL", "HT"},
	}
	fmt.Println("Mang a: ")
	printArray(a)
	fmt.Println("Mang b:")
	var b [3][2]string
	b[0][0] = "1"
	b[0][1] = "2"
	b[1][0] = "3"
	b[1][1] = "4"
	b[2][0] = "5"
	b[2][1] = "6"
	printArray(b)
}
