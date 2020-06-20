package delivery

import (
	"fmt"

	"github.com/neiln3121/dominos/game"
	"github.com/neiln3121/dominos/models"
)

// GameFunc - type for closure to be used with RepeatUntilNoError func
type GameFunc func(*game.Game, int) (bool, error)

// RepeatUntilNoError - Keeps reading input until a valid response us entered
func RepeatUntilNoError(fn GameFunc, current *game.Game, playerIndex int) bool {
	var err error
	var result bool
	// Keep going until no error
	for {
		result, err = fn(current, playerIndex)
		if err != nil {
			fmt.Printf("ERROR! - %v\n\n", err)
		} else {
			break
		}
	}
	return result
}

func choosePickup(current *game.Game, playerIndex int) (bool, error) {
	var picked int

	fmt.Printf("Pick a domino or 0 to return\n-> ")
	_, err := fmt.Scan(&picked)
	if err != nil {
		return false, err
	}
	if picked == 0 {
		return false, nil
	}
	err = current.PickDomino(playerIndex, picked)
	if err != nil {
		return false, err
	}

	return true, nil
}

func choosePlay(current *game.Game, playerIndex int) (bool, error) {
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

	err = current.PlayDomino(playerIndex, picked, end == 1)
	if err != nil {
		return false, err
	}

	return true, nil
}

func chooseUnpickedDominos(current *game.Game, player *models.Player) bool {
	fmt.Print(showUnpickedDominos(current.GetTable()))
	fmt.Printf("Player %d\n", player.ID)
	return RepeatUntilNoError(choosePickup, current, player.ID-1)
}

func chooseDominoToPlay(current *game.Game, player *models.Player) bool {
	fmt.Println(showBoard(current.GetBoard()))
	fmt.Println(showPlayerDominos(player))

	fmt.Printf("Player %d\n", player.ID)
	return RepeatUntilNoError(choosePlay, current, player.ID-1)
}
