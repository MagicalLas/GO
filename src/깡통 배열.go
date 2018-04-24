package main

import "fmt"

func main(){
	zero := []int{}
	var ze []int
	fmt.Println(zero, len(zero), cap(zero))
	fmt.Println(ze, len(ze), cap(ze))
	fmt.Println(nil)
	fmt.Println(zero==nil,ze==nil)
}