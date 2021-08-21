package main

import "fmt"

func returnFibonacci(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return returnFibonacci(n-2) + returnFibonacci(n-1)
}
func main() {
	fmt.Println("First 10 numbers of the fibonacci sequence:")
	for i := 0; i < 10; i++ {
		fmt.Println(returnFibonacci(i))
	}
}
