package main

func ret_two (a int, b int) (int, int){
	return b,a;
}

func main(){
a,v:=ret_two(5,9)
	println(a)
	println(v)
}