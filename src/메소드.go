package main

import(
	"fmt"
)
type Lab struct{
	a int
	b int
}
func (l *Lab) play(v int) int{
	return l.a+l.b+v
}
func main(){
	q := Lab{5,6}
	fmt.Println(q.play(9))
}