package main

import (
	"fmt"
)

func main() {
	buff := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(buff)
	for val := range buff {
		fmt.Println(val)
	}
}
