package main

import "fmt"

type trfle struct {
	X int
	Y string
	Z float32
}

func main() {
	str := trfle{Y: "las is good", X: 45, Z: 56.2}
	fmt.Println(str)
	fmt.Printf("%v", str)
}
