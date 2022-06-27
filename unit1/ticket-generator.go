package main

import (
	"fmt"
	"math/rand"
)

type ticket struct {
	spaceline string
	days      int
	tripType  string
	speed     int
	price     int
}

var spacelines = [3]string{"SpaceX", "Space Adventures", "Vrigin Galactic"}
var tripType = [2]string{"One-way", "Round-trip"}

const distance = 62100000
const basePrice = 36

func genTicket() ticket {
	var t ticket
	t.spaceline = spacelines[rand.Intn(3)]
	t.tripType = tripType[rand.Intn(2)]
	t.speed = rand.Intn(14) + 16
	t.days = distance / t.speed / 24 / 3600
	var extraPrice = float32(t.speed) / 30 * 14
	if t.tripType == tripType[0] {
		t.price = basePrice + int(extraPrice)
	} else {
		t.price = (basePrice + int(extraPrice)) * 2
	}
	return t
}

func main() {
	tickets := make([]ticket, 0, 10)

	ticketsSlice := tickets[:]

	for i := 0; i < 10; i++ {
		var t ticket = genTicket()
		ticketsSlice = append(ticketsSlice, t)
	}

	fmt.Printf("%-17v%4v %-11v%5v  \n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Print("======================================\n")

	for _, t := range ticketsSlice {
		fmt.Printf("%-17v%4v %-11v%5v  \n", t.spaceline, t.days, t.tripType, t.price)
	}
	// fmt.Print(len(tickets))

}
