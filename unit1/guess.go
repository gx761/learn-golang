package main

import (
	"fmt"
	"math/rand"
)

func main(){

	myNumber:=50

	for  found:= false; found != true; {
		random := rand.Intn(100)

		if(random == myNumber) {
			found = true;
			fmt.Println(random, "is exactly what I want")
		} else if(random < myNumber){
			fmt.Println(random, "is smaller than my number")
		} else {
			fmt.Println(random, "is larger than my number")
		}
	}

}