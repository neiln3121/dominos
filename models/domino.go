package models

import "fmt"

// DOTS - max number of dots on each domino half
const DOTS int = 6

// Domino - stores the game play piece
type Domino struct {
	Half          [2]int
	IsPicked      bool
	PlayedFlipped bool
}

// Total - total number of dots on a domino
func (d *Domino) Total() int {
	return d.Half[0] + d.Half[1]
}

// Set - sets the number of dots on each domino half
func (d *Domino) Set(head, tail int) error {
	if head > DOTS || head < 0 {
		return fmt.Errorf("Invalid value for domino: %d", head)
	}
	if tail > DOTS || tail < 0 {
		return fmt.Errorf("Invalid value for domino: %d", tail)
	}
	d.Half[0] = head
	d.Half[1] = tail

	return nil
}