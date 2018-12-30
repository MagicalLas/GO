package main

import "fmt"

func main() {
	arr := []int{2, 2, 5, 4, 5, 36, 37, 28, 19}
	for i, t := range arr {
		fmt.Printf("%v %v \n", i, t)
	}
}
