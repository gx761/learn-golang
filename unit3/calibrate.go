package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	offset := 5
	sensor := calibrate(fakeSensor, kelvin(offset))
	fmt.Println(sensor())
	fmt.Println(sensor())
	fmt.Println(sensor())
}
