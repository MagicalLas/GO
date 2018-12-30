package main

import "fmt"

type point struct {
	X int
	Y int
}

func main() {

	a := point{5, 6}
	fmt.Println(point{2, 3})

	fmt.Println(a)
}
