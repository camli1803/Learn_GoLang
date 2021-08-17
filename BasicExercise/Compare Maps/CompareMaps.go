package main

import (
	"fmt"
)

func Comparetwomaps(map1, map2 map[string]int) {
	check := true
	if len(map1) == len(map2) { //check chiều dài đảm bảo mọi key trong map1 đều được kiểm tra trong map2 và ngược lại
		for key, value := range map1 {
			valuemap2, ok := map2[key] //check xem map2 có key đó giống map1 không, ok là biến check.
			if ok == true {
				if valuemap2 != value {
					check = false
				}
			} else {
				check = false
			}
		}
	} else {
		check = false
	}
	if check == true {
		fmt.Println("Map giống nhau!")
	} else {
		fmt.Println("Map khác nhau!")
	}
}

func main() {
	map1 := make(map[string]int)
	map1["CL"] = 1803
	map1["MC"] = 1710

	map2 := make(map[string]int)
	map2["CL"] = 1803
	map2["MC"] = 1710

	Comparetwomaps(map1, map2)

}
