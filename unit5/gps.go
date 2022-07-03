package main

import (
	"fmt"
	"math"
)

type location struct {
	name string
	lat  float64
	long float64
}

type world struct {
	radius float64
}

type gps struct {
	current     location
	destination location
	world
}

type coordinate struct {
	d, m, s float64
	h       rune
}

type rover struct {
	gps
}

func (l location) description() {
	fmt.Println("name: ", l.name, "latitude: ", l.lat, "longtitude: ", l.long)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() {
	fmt.Println(g.distance(), "kilometers are left")
}

func main() {
	w := world{
		6371.0,
	}
	current := location{"Bradbury Landing", -4.5895, 137.4417}
	destination := location{"Elysium Planitia", 4.5, 135.9}
	g := gps{
		current,
		destination,
		w,
	}
	curiosity := rover{g}
	curiosity.message()
}
