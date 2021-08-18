//geometry.go
package main

import (
	"fmt"

	"github.com/camli1803/Learn_Golang/geometry/rectangle" //import package from github by module.

	"log"
	_ "math" //Blank identifier
)

var rectLen, rectWidth float64 = 6, 7

func init() {
	fmt.Println("main package initialized!")
	if rectLen < 0 {
		log.Fatal("Error: Lenght is less than 0")
	}
	if rectWidth < 0 {
		log.Fatal("Error: Width is less than 0")
	}

}
func main() {
	fmt.Println("Geometrical shape properties")

	/*Hàm Area của package rectangle được sử dụng
	 */
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Hàm Diagonal của package rectangle được sử dụng
	 */
	fmt.Printf("Diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
