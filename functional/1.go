package main

func main() {
	a := func(a, b int) (result int) {
		result = a + b
		return
	}
	print(a(3, 4))
}
