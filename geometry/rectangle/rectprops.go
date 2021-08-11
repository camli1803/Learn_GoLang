//rectprops.go
package rectangle

import (
	"math"

	"fmt"
)

func init() {
	fmt.Println("rectangle package initialized!")
}

func Area(len, wid float64) float64 { // viết hoa chữ cái đầu tiên thì mới có thể truy cập được từ bên ngoài pkg
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
