package main

import (
	"fmt"
)

func main() {
	var buff = [...]int{1, 2, 3, 4, 5}
	fmt.Println(buff)
	var buff_1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(buff_1)
	var buff_2 = [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(buff_2)
}
