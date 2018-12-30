package main

import (
	"fmt"
)

type two struct {
	X int
	Y int
}

func main() {
	pointer := new(two)
	fmt.Println(pointer)
	pointer.X = 5
	pointer.Y = 15
	fmt.Println(pointer)
}
