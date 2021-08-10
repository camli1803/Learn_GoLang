//geometry.go
package main

import (
	"fmt"
	"geometry/rectangle" //importing package tùy chỉnh
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")

	/*Hàm Area của package rectangle được sử dụng
	 */
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Hàm Diagonal của package rectangle được sử dụng
	 */
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
