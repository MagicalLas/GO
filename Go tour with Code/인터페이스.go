package main

import ("fmt")

type Las interface{
	l(int) int
}

func main(){
	var a Las
	a = 5

	a.l(5)
}

type Luv int

func (l Luv) l (v int) int{
	fmt.Println(v)
	return 5
}