package main

import "fmt"

type Planets []string

func (p Planets) Terraform() Planets {
	for i, v := range p {
		p[i] = "New " + v
	}
	return p
}

func main() {
	var planets Planets = []string{"Mars", "Uranus", "Neptune"}
	planets.Terraform()
	fmt.Println(planets)
}
