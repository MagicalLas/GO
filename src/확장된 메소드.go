package main

import "fmt"

type las int

func (f las) play(a int, q int) float64 {
	return float64(a+q)
}

func main() {
	f := las(3)

	fmt.Println(f.play(2,5))
}