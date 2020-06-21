package models

import (
	"errors"
)

// DOTS - max number of dots on each domino half
const DOTS int = 6

// Domino - stores the game play piece
type Domino struct {
	IsPicked      bool
	half          [2]int
	playedFlipped bool
}

// GetDots - get the dots in order according to how domino was played
func (d *Domino) GetDots() (int, int) {
	if d.playedFlipped {
		return d.half[1], d.half[0]
	}
	return d.half[0], d.half[1]
}

// Total - total number of dots on a domino
func (d *Domino) Total() int {
	return d.half[0] + d.half[1]
}

// Set - sets the number of dots on each domino half
func (d *Domino) Set(head, tail int) error {
	if head > DOTS || head < 0 || tail > DOTS || tail < 0 {
		return errors.New("Invalid value for domino half")
	}

	d.half[0] = head
	d.half[1] = tail

	return nil
}
