package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 20; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("Đây là kết quả hàm tự viết!")
	fmt.Println(Sqrt1(2))
	fmt.Println("Đây là kết quả từ thư viện Math!")
	fmt.Println(math.Sqrt(2))
}
