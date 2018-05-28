package main

func add_x(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

func main(){
	println(add_x(3)(4))
}