package main

import (
	"fmt"
	"strconv"
)

func main() {
	defaultSlice := []string{"1", "2"}
	for i := 0; i < 20; i++ {
		lastNumber, _ := strconv.ParseInt(defaultSlice[len(defaultSlice)-1], 0, 32)

		defaultSlice = append(defaultSlice, string(lastNumber))

		fmt.Printf("length is %v, capacity is %v \n", len(defaultSlice), cap(defaultSlice))

	}

}
