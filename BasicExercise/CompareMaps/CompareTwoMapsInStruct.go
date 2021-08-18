package main

import (
	"github.com/camli1803/Learn_Golang/BasicExercise/CompareMaps/comparemapsfunction"
)

type image struct {
	data map[int]int
}

func main() {
	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	comparemapsfunction.PrintNotice(image1.data, image2.data)

}
