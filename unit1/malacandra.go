package main

import (
	"fmt"
)

func main() {

	duration:=28.;
	distance:=56000000.;

	speed:= distance/duration/24

	fmt.Printf("%v, %T", speed, speed)

}