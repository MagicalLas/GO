package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["las"] = 5
	m["la"] = 2

	fmt.Println(m)

	var mm = map[string]int{
		"a": 5,
		"B": 8,
	}
	fmt.Println(mm)

}
