package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(arr)
	fmt.Println("[n]", arr[1])
	fmt.Println("[:n]", arr[:5])
	fmt.Println("[n:]", arr[5:])
	fmt.Println("[n:m]", arr[5:8])
}
