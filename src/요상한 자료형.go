package main

func main(){
	var(
		//허수???????? 그리고 이게 가능한건가... 대단해..
		ex complex128 = 300000000000000<<400
		a uintptr = 300000000000000<<15
	)
	println(ex, a)
}