package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	uni := make([][]bool, width)

	for i := range uni {
		uni[i] = make([]bool, height)
	}
	return uni
}

func (u Universe) Show() {
	for _, v := range u {
		for _, v2 := range v {
			if v2 == true {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	for _, v := range u {
		for i := range v {
			random := rand.Intn(100)
			if random <= 25 {
				v[i] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	var finalX, finalY int
	if x >= 0 && x < width {
		finalX = x
	} else if x >= width {
		finalX = x % width
	} else {
		finalX = x%width + width
	}
	if y >= 0 && y < height {
		finalY = y
	} else if y >= height {
		finalY = y % height
	} else {
		finalY = y%height + height
	}
	return u[finalX][finalY]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if u.Alive(i, j) == true {
				count++
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	count := u.Neighbors(x, y)
	if u.Alive(x, y) == true {
		if count < 2 || count > 3 {
			return false
		}
		return true
	} else {
		if count == 3 {
			return true
		}
		return false
	}
}

func Step(a, b Universe) {
	for i, v := range b {
		for j := range v {
			// fmt.Print(i, j)
			b[i][j] = a.Next(i, j)
		}
	}
}

func main() {
	current, next := NewUniverse(), NewUniverse()
	fmt.Printf("%T", current)
	current.Seed()
	// current.Show()

	for true {
		Step(current, next)
		fmt.Print("\033[H")
		// next.Show()
		current, next = next, current
		time.Sleep(time.Second * 2)
	}

}
