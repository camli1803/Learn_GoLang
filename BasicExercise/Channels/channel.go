package main

import "fmt"

func digits(a int, digitchannel chan int) {
	for a != 0 {
		digit := a % 10
		digitchannel <- digit
		a /= 10
	}
	close(digitchannel)
}

func calSquare(a int, squarechannel chan int) {
	sum := 0
	digitchannel := make(chan int)
	go digits(a, digitchannel)
	for digit := range digitchannel {
		sum += digit * digit
	}
	squarechannel <- sum

}

func calcubes(a int, cubeschannel chan int) {
	sum := 0
	digitchannel := make(chan int)
	go digits(a, digitchannel)
	for digit := range digitchannel {
		sum += digit * digit * digit
	}
	cubeschannel <- sum
}

func main() {
	a := 534
	squarechannel := make(chan int)
	cubeschannel := make(chan int)
	go calSquare(a, squarechannel)
	go calcubes(a, cubeschannel)
	fmt.Println("Square:", <-squarechannel, "Cubes:", <-cubeschannel)
}
