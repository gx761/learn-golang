package main

import (
	"fmt"
	"math"
	"strings"
)

type getData func(v float64) float64
type cal func() float64
type celsius float64
type fahrenheit float64

func main() {

	var test = func(t float64) float64 {
		return t + 5.0
	}

	drawTable(test)
}

func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit((c * 9 / 5) + 32)
}

func calibrate(start float64, get getData) cal {
	remain := start

	return func() float64 {
		remain = get(remain)
		return remain
	}
}

func drawTable(get getData) {
	g := calibrate(-45.0, get)
	b := []byte{0xE2, 0x84, 0x83}

	fmt.Println(strings.Repeat("=", 23))
	fmt.Println("|", string(b), strings.Repeat(" ", 6), "|", "℉", strings.Repeat(" ", 6), "|")
	fmt.Println(strings.Repeat("=", 23))

	c := g()

	for c <= 100 {
		// fmt.Println("|", c, strings.Repeat(" ", 6), "|", c, strings.Repeat(" ", 6), "|")
		fmt.Printf("| %-8v | %-8v |\n", math.Floor(c*100)/100, celsiusToFahrenheit(celsius(c)))
		c = g()
	}

	fmt.Println(strings.Repeat("=", 23))
}
