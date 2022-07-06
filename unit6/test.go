package main

import "fmt"

func main() {
	var v interface{}
	var x int
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	fmt.Printf("%T %v \n", x, x)
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)
}
