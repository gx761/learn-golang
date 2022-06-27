package main

import (
	"fmt"
	"math/rand"
)

func isLeap(year int) bool{
	if(year % 4 == 0 || year % 400 ==0) {
		return true
	}
	return false
}

func genDate() {
	year := rand.Intn(2022) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31
	era := "AD"
switch month {
	case 2:
		if(isLeap(year)) {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4,6,9,11:
		daysInMonth = 30
}
	day:=rand.Intn(daysInMonth) +1
	fmt.Println(era, year, month, day)
}

func main() {

	for i:=0;i<10;i++ {
		genDate()
	}
}	