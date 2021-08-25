package main

import (
	"fmt"
)

func digits(a int, digitchan chan int) {
	for a != 0 {
		digitnum := a % 10
		digitchan <- digitnum
		a /= 10
	}
	close(digitchan)
}
func calSquare(a int, squarechan chan int) {
	digital := make(chan int)
	go digits(a, digital)
	sum := 0
	for digit := range digital {
		sum += digit * digit
	}
	squarechan <- sum

}
func calCubes(a int, cubeschan chan int) {
	digital := make(chan int)
	go digits(a, digital)
	sum := 0
	for digit := range digital {
		sum += digit * digit * digit
	}
	cubeschan <- sum
}
func main() {
	a := 589
	squarechan := make(chan int)
	cubeschan := make(chan int)
	go calSquare(a, squarechan)
	go calCubes(a, cubeschan)
	squarechanint := <-squarechan
	cubeschanint := <-cubeschan
	fmt.Println(squarechanint + cubeschanint)
}

/*package main

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
} */
