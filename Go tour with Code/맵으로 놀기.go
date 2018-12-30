package main

import "fmt"

func main() {
	M := make(map[string]int)
	M["a"] = 2
	M["q"] = 5
	fmt.Println(M["a"])

	M["a"] = 50
	fmt.Println(M["a"])

	delete(M, "a")
	fmt.Println(M["a"])

	v, o := M["a"]
	fmt.Println(v, o)

	z, x := M["q"]
	fmt.Println(z, x)
}
