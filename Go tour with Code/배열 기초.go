package main

func main(){
	pointer := []int{1,2,3,4,5}
	for i := 0; i < 5; i++ {
		println(pointer[i])
	}
	println(pointer)
}