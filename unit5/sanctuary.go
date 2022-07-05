package main

import "math/rand"

type dog struct {
	name string
}

type cat struct {
	name string
}

type animal interface {
	eat() string
	move() string
	String() string
}

type santuary []animal

func (d dog) String() string {
	return d.name
}

func (d dog) move() string {
	i := rand.Intn(2)

	switch i {
	case 0:
		return "get up"
	case 1:
		return "get down"
	default:
		return "get down"
	}
}

func (d dog) eat() string {
	i := rand.Intn(2)
	switch i {
	case 0:
		return "eat shit"
	case 1:
		return "eat cat food"
	default:
		return "eat cat food"
	}
}

func (d cat) String() string {
	return d.name
}

func (d cat) move() string {
	i := rand.Intn(2)

	switch i {
	case 0:
		return "get up"
	case 1:
		return "get down"
	default:
		return "get down"
	}
}

func (d cat) eat() string {
	i := rand.Intn(2)
	switch i {
	case 0:
		return "eat shit"
	case 1:
		return "eat cat food"
	default:
		return "eat cat food"
	}
}

func main() {
	dog := dog{"doudou"}
	cat := cat{"xiaoming"}

	animals := []animal{dog, cat}

	sols := 3
	dayLong := 24

}
