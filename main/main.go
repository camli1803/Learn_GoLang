package main

import (
	"fmt"
)

func main() {
	price := 12000
	sale := 2000
	no := 12
	totalPrice, _ := caculateBill(price, sale, no)
	fmt.Println(totalPrice)
}
func caculateBill(price int, sale int, no int) (int, int) {
	totalPrice := price * no
	rePrice := price - sale
	return totalPrice, rePrice
}
