package main

func retTwo(a int, b int) (int, int) {
	return b, a
}

func main() {
	a, v := retTwo(5, 9)
	println(a)
	println(v)
}
