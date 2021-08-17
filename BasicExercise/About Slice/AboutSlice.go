package main

import (
	"fmt"
)

func main() {
	FatherArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	ChildSlice := FatherArray[3:4]
	fmt.Println("Before re-slice:")
	fmt.Println(ChildSlice)
	fmt.Println("Length:", len(ChildSlice), "Capacity:", cap(ChildSlice))
	ChildSlice = ChildSlice[:cap(ChildSlice)]
	fmt.Println("After re-slice:")
	fmt.Println(ChildSlice)
	fmt.Println("Length:", len(ChildSlice), "Capacity:", cap(ChildSlice))
}
