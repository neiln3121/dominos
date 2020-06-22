package delivery

import (
	"fmt"

	"github.com/neiln3121/dominos/game"
)

// SelectOption - Cycle through options for current game player
func SelectOption(current *game.Game) {
	var picked int
	var success bool

	for !success {
		fmt.Println(showBoard(current.GetBoard()))
		player := current.GetCurrentPlayer()
		fmt.Printf("Player %s\n", player.Name)
		fmt.Println(showPlayerDominos(player))

		fmt.Print("Pick a option\n1: Play\n2: Pick up\n-> ")
		_, err := fmt.Scan(&picked)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		if picked < 1 || picked > 2 {
			fmt.Printf("Invalid option: must be 1 or 2\n\n")
			continue
		}
		if picked == 1 {
			success = chooseDominoToPlay(current)
		}
		if picked == 2 {
			success = chooseUnpickedDominos(current)
		}
	}
	fmt.Println("\nSuccess!")
	fmt.Println(showBreak())
}

// SelectPlayerNames - Set the player names
func SelectPlayerNames(current *game.Game) error {
	for i, player := range current.GetPlayers() {
		fmt.Printf("Pick a name for Player %d-> ", i+1)
		_, err := fmt.Scan(&player.Name)
		if err != nil {
			return err
		}
	}
	fmt.Println("\nDone!")
	fmt.Println(showBreak())
	return nil
}

// SelectInitialDominos - Set up all players by picking up their initial dominos
func SelectInitialDominos(current *game.Game) {
	var success bool

	for !success {
		success = chooseUnpickedDominos(current)
		if !success {
			fmt.Print("You must pick a domino during setup!\n\n")
		}
	}
	fmt.Println(showBreak())
}

func chooseUnpickedDominos(current *game.Game) bool {
	fmt.Print(showUnpickedDominos(current.GetTable()))
	fmt.Printf("Player %s\n", current.GetCurrentPlayer().Name)
	return repeatUntilNoError(vaidatePickup, current)
}

func chooseDominoToPlay(current *game.Game) bool {
	fmt.Println(showBoard(current.GetBoard()))
	fmt.Println(showPlayerDominos(current.GetCurrentPlayer()))

	fmt.Printf("Player %s\n", current.GetCurrentPlayer().Name)
	return repeatUntilNoError(validatePlay, current)
}
