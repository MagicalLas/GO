package main

func cal(function func(int, int) int, a int, b int) {
	println(function(a, b))
}
func mal(a int, b int) int {
	return a * b
}
func main() {
	cal(mal, 5, 6)
}
