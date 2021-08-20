package main

import "fmt"

func detmatrix(a [3][3]int) (detmatrix int) {
	detmatrix = a[0][0]*a[1][1]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[1][0]*a[2][1] - a[0][2]*a[1][1]*a[2][0] - a[0][1]*a[1][0]*a[2][2] - a[0][0]*a[1][2]*a[2][1]
	return
}

func main() {
	var matrix [3][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 2
	matrix[1][1] = 3
	matrix[1][2] = 1
	matrix[2][0] = 3
	matrix[2][1] = 1
	matrix[2][2] = 2
	fmt.Println("Det của ma trận là:", detmatrix(matrix))

}
