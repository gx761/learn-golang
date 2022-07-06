package main

import "fmt"

type item string

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup() {
	i := item("sword")
	c.leftHand = &i
	fmt.Println("pick up a ", i)
}

func (i *character) give(e *character) {
	e.leftHand = i.leftHand
	i.leftHand = nil
}

func main() {
	au := &character{"author", nil}
	knight := &character{"knight", nil}
	au.pickup()
	au.give(knight)
	fmt.Printf("%+v\n", au)
	fmt.Printf("%+v\n", knight)
}
