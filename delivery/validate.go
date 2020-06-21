package delivery

import (
	"fmt"

	"github.com/neiln3121/dominos/game"
)

// gameFunc - type for closure to be used with RepeatUntilNoError func
type gameFunc func(*game.Game) (bool, error)

// repeatUntilNoError - Keeps reading input until a valid response us entered
func repeatUntilNoError(fn gameFunc, current *game.Game) bool {
	var err error
	var result bool
	// Keep going until no error
	for {
		result, err = fn(current)
		if err != nil {
			fmt.Printf("ERROR! - %v\n\n", err)
		} else {
			break
		}
	}
	return result
}

func vaidatePickup(current *game.Game) (bool, error) {
	var picked int

	fmt.Printf("Pick a domino or 0 to return\n-> ")
	_, err := fmt.Scan(&picked)
	if err != nil {
		return false, err
	}
	if picked == 0 {
		return false, nil
	}
	err = current.PickDomino(picked)
	if err != nil {
		return false, err
	}

	return true, nil
}

func validatePlay(current *game.Game) (bool, error) {
	var picked int
	var end int

	fmt.Print("Play a domino\n-> ")
	_, err := fmt.Scan(&picked)
	if err != nil {
		return false, err
	}
	fmt.Print("Play at either the left(1) or right(2) or 0 to return\n-> ")
	_, err = fmt.Scan(&end)
	if err != nil {
		return false, err
	}
	if end == 0 {
		return false, nil
	}
	if end < 1 || end > 2 {
		return false, fmt.Errorf("Invalid end: must be 1 or 2")
	}

	err = current.PlayDomino(picked, end == 1)
	if err != nil {
		return false, err
	}

	return true, nil
}
