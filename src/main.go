package main

import (
	"fmt"
)

func Mycal(a int, b int) (int, int) {
	return (a + b), (a - b)
}

func main() {
	temp_1, _ := Mycal(10, 10)

	for {
		if temp_1 > 30 {
			break
		}
		fmt.Println(temp_1)
		temp_1 = temp_1 + 1
		fmt.Println("aaaaa")
	}
}
