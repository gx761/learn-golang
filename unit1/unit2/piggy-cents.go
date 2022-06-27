package main

import (
	"fmt"
	"math/rand"
)


func main() {
	balance:=20.00
	deposits:= []int{5, 10, 25}

	for balance >=0 {
		deposit:= deposits[rand.Intn(2)]
		balance = balance - float64(deposit)/100;
		fmt.Printf("Amount of $%4.2f is deposited, now balance: $%4.2f\n",float64(deposit)/100, balance)
	}
}