package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	f := celsiusToFahrenheit(c)
	return f
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "o K is ", celsius, "o C")
}
