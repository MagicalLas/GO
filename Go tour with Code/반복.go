package main

import "fmt"

var pow = []int{ 10,9,8,7,6,5,4,3,2,1}

func main() {
	for i, v := range pow {
		fmt.Printf("%d = %d\n", i, v)
	}
}
