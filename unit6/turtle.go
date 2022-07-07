package main

import "fmt"

type turtle struct {
	x int
	y int
}

func (a *turtle) up() {
	a.y--
}

func (a *turtle) down() {
	a.y++
}

func (a *turtle) right() {
	a.x++
}

func (a *turtle) left() {
	a.x--
}

func main() {
	t := turtle{15, 15}
	t.up()
	fmt.Println(t)
	t.right()
	fmt.Println(t)
	t.down()
	fmt.Println(t)
	t.left()
	fmt.Println(t)
}
