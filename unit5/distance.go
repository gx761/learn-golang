package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}
type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat           coordinate
	long          coordinate
	site          string
	roverOrLander string
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat.decimal()))
	s2, c2 := math.Sincos(rad(p2.lat.decimal()))
	clong := math.Cos(rad(p1.long.decimal() - p2.long.decimal()))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func newLocation(lat coordinate, long coordinate, roverOrLander string, site string) location {
	return location{
		lat,
		long,
		site,
		roverOrLander,
	}
}

func getLocations() []location {

	locations := []location{
		{
			roverOrLander: "Spirit",
			site:          "Columbia Memorial Station",
			lat: coordinate{
				14, 34, 5.2, 'S',
			},
			long: coordinate{
				175, 28, 21.5, 'E',
			},
		},
		{
			roverOrLander: "Opportunity",
			site:          "Challenger Memorial Station",
			lat: coordinate{
				1, 56, 46.3, 'S',
			},
			long: coordinate{
				354, 28, 24.2, 'E',
			},
		},
		{
			roverOrLander: "Curiosity",
			site:          "Bradbury Landing",
			lat: coordinate{
				4, 35, 22.2, 'S',
			},
			long: coordinate{
				137, 26, 30.1, 'E',
			},
		},
		{
			roverOrLander: "InSight",
			site:          "Elysium Planitia",
			lat: coordinate{
				4, 30, 0.0, 'N',
			},
			long: coordinate{
				135, 54, 0, 'E',
			},
		},
	}
	return locations
}

type pair struct {
	start    location
	end      location
	distance float64
}

func calculate(w world, locations []location) (pair, pair) {
	var shortest pair
	var fathest pair
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			d := w.distance(locations[i], locations[j])
			if d < shortest.distance || shortest.distance == 0 {
				shortest.start = locations[i]
				shortest.end = locations[j]
				shortest.distance = d
			}
			if d > fathest.distance || fathest.distance == 0 {
				fathest.start = locations[i]
				fathest.end = locations[j]
				fathest.distance = d
			}
		}
	}
	return shortest, fathest
}

func main() {
	worlds := map[string]world{
		"Mercury": {2439.7},
		"Venus":   {6051.8},
		"Earth":   {6371.0},
		"Mars":    {3389.5},
		"Jupiter": {69911},
		"Saturn":  {58232},
		"Uranus":  {25362},
		"Neptune": {24622},
	}
	locations := getLocations()
	shortest, fathest := calculate(worlds["Earth"], locations)
	fmt.Printf("%+v\n", shortest)
	fmt.Printf("%+v\n", fathest)

}
