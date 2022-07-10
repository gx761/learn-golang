package main

import (
	"errors"
	"fmt"
)

type Sudoku [9][9]int8

const (
	rows = 9
	cols = 9
)

var (
	ErrBounds     = errors.New("out of bounds")
	ErrInRow      = errors.New("same value already exist on the same row")
	ErrDigit      = errors.New("digit is not valid")
	ErrInRegion   = errors.New("same value exist in the same region")
	ErrInCol      = errors.New("same value already exist on the same column")
	ErrFixedDigit = errors.New("original value cannot be overidden")
)

func NewSudoku(ini [9][9]int8) Sudoku {
	return Sudoku(ini)
}

func main() {
	s := NewSudoku([rows][cols]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	original := s

	err := s.setDigit(1, 2, 10, original)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("set digit successfully")
	}
}

func (s *Sudoku) checkInbounds(x int, y int) bool {
	if x >= rows || x < 0 || y >= cols || y < 0 {
		return false
	}
	return true
}

func (s *Sudoku) checkFixed(x int, y int, orignal *Sudoku) bool {
	if orignal[x][y] != 0 {
		return false
	}
	return true
}

func (s *Sudoku) checkHorizontal(x int, y int, v int8) bool {
	for j := 0; j < cols; j++ { //check horizontal value
		if j == y {
			continue
		}
		if s[x][j] == v {
			return false
		}
	}
	return true
}

func (s *Sudoku) checkVertical(x int, y int, v int8) bool {
	for i := 0; i < rows; i++ { //check vertical value
		if i == y {
			continue
		}
		if s[i][y] == v {
			return false
		}
	}
	return true
}

func (s *Sudoku) checkSection(x int, y int, v int8) bool {
	startRow := x / 3 * 3
	startCol := y / 3 * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if i == x && j == y {
				continue
			}
			if s[x][y] == v {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) setDigit(x int, y int, v int8, original Sudoku) error {
	if !s.checkInbounds(x, y) {
		return ErrBounds
	}

	if !s.checkFixed(x, y, &original) {
		return ErrFixedDigit
	}

	if !s.checkHorizontal(x, y, v) {
		return ErrInRow
	}

	if !s.checkVertical(x, y, v) {
		return ErrInCol
	}

	if !s.checkSection(x, y, v) {
		return ErrInRegion
	}

	return nil

}
