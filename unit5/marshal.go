package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

func (c coordinate) MarshalJSON() ([]byte, error) {
	var detailed struct {
		Decimal    float64
		Dms        string
		Degrees    float64
		Minutes    float64
		Seconds    float64
		Hemisphere byte
	}
	detailed.Decimal = c.decimal()
	detailed.Dms = strconv.FormatFloat(c.d, 'f', -1, 64) + `°` + strconv.FormatFloat(c.m, 'f', -1, 64) + `'` + strconv.FormatFloat(c.s, 'f', -1, 64) + `" ` + string(c.h)
	detailed.Degrees = c.d
	detailed.Minutes = c.m
	detailed.Seconds = c.s
	detailed.Hemisphere = byte(c.h)

	return json.Marshal(detailed)
}

func main() {
	c := coordinate{135, 54, 0.0, 'E'}
	bytes, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(bytes))
}
