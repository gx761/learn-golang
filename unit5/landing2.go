package main

import (
	"fmt"
	"os"
)

type location struct {
	Lat           coordinate
	Long          coordinate
	Site          string
	RoverOrLander string
}
type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func getLocations() []location {

	locations := []location{
		{
			RoverOrLander: "Spirit", Site: "Columbia Memorial Station",
			Lat: coordinate{
				14, 34, 5.2, 'S',
			},
			Long: coordinate{
				175, 28, 21.5, 'E',
			},
		},
		{
			RoverOrLander: "Opportunity", Site: "Challenger Memorial Station",
			Lat: coordinate{
				1, 56, 46.3, 'S',
			},
			Long: coordinate{
				354, 28, 24.2, 'E',
			},
		},
		{
			RoverOrLander: "Curiosity", Site: "Bradbury Landing",
			Lat: coordinate{
				4, 35, 22.2, 'S',
			},
			Long: coordinate{
				137, 26, 30.1, 'E',
			},
		},
		{
			RoverOrLander: "InSight", Site: "Elysium Planitia",
			Lat: coordinate{
				4, 30, 0.0, 'N',
			},
			Long: coordinate{
				135, 54, 0, 'E',
			},
		},
	}
	return locations
}

func main() {
	locations := getLocations()

	fmt.Printf("%-25v%-35v%15v%15v\n", "Rover or lander", "Landing site", "Latitude", "Longitude")
	for _, v := range locations {
		fmt.Printf("%-25v%-35v%15.2f%15.2f\n", v.RoverOrLander, v.Site, v.Lat.decimal(), v.Long.decimal())
	}

}
func exitOnError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
