package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

type RoverDriver struct {
	commandc chan command
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Microsecond
	nextMove := time.After(updateInterval)
	speed := 1
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case start:
				speed = 1
				log.Printf("start driving")
			case stop:
				speed = 0
				log.Printf("stop driving")
			case right:
				direction = image.Point{X: -direction.X, Y: direction.Y}
			case left:
				direction = image.Point{X: direction.X, Y: -direction.Y}
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction.Mul(speed))
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}
func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
}
