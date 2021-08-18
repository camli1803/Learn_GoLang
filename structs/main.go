package main

import (
	"fmt"

	"github.com/camli1803/Learn_Golang/structs/computer"
)

func main() {
	var person computer.Person
	person.Name = "Cam Li"
	person.Age = 21
	person.City = "Hà Nội"
	person.State = "Tảo Khê"
	fmt.Println(person)
}
