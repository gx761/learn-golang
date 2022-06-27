package main

import (
	"fmt"
	"math/rand"
)


func main() {
	balance:=20.00
	deposits:= []float64{0.05, 0.10, 0.25}

	for balance >=0 {
		deposit:= deposits[rand.Intn(2)]
		balance = balance - deposit;
		fmt.Printf("Amount of $%4.2f is deposited, now balance: $%4.2f\n",deposit, balance)
	}
}