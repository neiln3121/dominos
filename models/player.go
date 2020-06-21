package models

import (
	"errors"
)

// STARTING_DOMINOS - number of dominos a player should start with
const STARTING_DOMINOS int = 7

// Player - stores detaisl for a game participant
type Player struct {
	Name    string
	dominos []*Domino
}

// AddDomino - add a domino to a player
func (p *Player) AddDomino(domino *Domino) error {
	if domino == nil {
		return errors.New("Not a valid domino")
	}
	p.dominos = append(p.dominos, domino)

	return nil
}

// RemoveDomino - remove a domino to a player
func (p *Player) RemoveDomino(index int) error {
	if !p.HasDomino(index) {
		return errors.New("Player does not have that domino")
	}
	p.dominos = append(p.dominos[:index], p.dominos[index+1:]...)

	return nil
}

// GetHighestDouble - get the highest double held, used to work out starting player
func (p *Player) GetHighestDouble() (highest, index int) {
	highest = -1
	for i, domino := range p.dominos {
		// Must be a double
		if domino.half[0] == domino.half[1] {
			if domino.half[0] > highest {
				highest = domino.half[0]
				index = i
			}
		}
	}
	return
}

// CanProceed - calculates if the player can continue given two board ends
func (p *Player) CanProceed(head, tail int) bool {
	for _, domino := range p.GetDominos() {
		if domino.half[0] == head || domino.half[0] == tail ||
			domino.half[1] == head || domino.half[1] == tail {
			return true
		}
	}
	return false
}

// Get - get a player's domino at specified index
func (p *Player) Get(index int) (*Domino, error) {
	if !p.HasDomino(index) {
		return nil, errors.New("Player does not have that domino")
	}

	return p.dominos[index], nil
}

// GetDominos - get all the player's dominos
func (p *Player) GetDominos() []*Domino {
	return p.dominos
}

// HasStartingDominos - check if player has enough dominos to start the game
func (p *Player) HasStartingDominos() bool {
	return p.DominoCount() == STARTING_DOMINOS
}

// DominoCount - number of dominos the player currently has
func (p *Player) DominoCount() int {
	return len(p.dominos)
}

// TotalDots - total dots on each domino held by the player, used to calculate the winner
func (p *Player) TotalDots() int {
	sum := 0
	for _, domino := range p.dominos {
		sum += domino.Total()
	}
	return sum
}

// HasDomino
func (p *Player) HasDomino(index int) bool {
	if index < 0 || index >= p.DominoCount() {
		return false
	}
	return true
}

// NewPlayer - new instance
func NewPlayer() *Player {
	dominos := new([]*Domino)
	return &Player{dominos: *dominos}
}
