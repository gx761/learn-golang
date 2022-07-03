package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string  `json:"name"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main() {
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	bytes, err := json.MarshalIndent(locations, "", "  ")
	// bytes, err := json.Marshal(locations)
	exitOnError(err)

	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
