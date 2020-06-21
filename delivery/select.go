package delivery

import (
	"fmt"

	"github.com/neiln3121/dominos/game"
	"github.com/neiln3121/dominos/models"
)

// SelectOption - Cycle through options for current game player
func SelectOption(current *game.Game) error {
	var picked int
	var err error
	var success bool

	for !success {
		fmt.Println(showBoard(current.GetBoard()))
		player := current.GetCurrentPlayer()
		fmt.Println(showPlayerDominos(player))
		fmt.Printf("Player %d\n", player.ID)

		fmt.Print("Pick a option\n1: Play\n2: Pick up\n-> ")
		_, err = fmt.Scan(&picked)
		if err != nil {
			return err
		}
		if picked < 1 || picked > 2 {
			fmt.Printf("Invalid option: must be 1 or 2\n\n")
			continue
		}
		if picked == 1 {
			success = chooseDominoToPlay(current, player)
		}
		if picked == 2 {
			success = chooseUnpickedDominos(current, player)
		}
	}
	fmt.Println("\nSuccess!")
	fmt.Println(showBreak())
	return nil
}

// SelectInitialDominos - Get all players by picking up their initial dominos
func SelectInitialDominos(current *game.Game) {
	var allPlayersReady int
	// Keep going until all players have the right amount of dominos
	for allPlayersReady < len(current.GetPlayers()) {
		// Set back to zero - every player needs the required amount of dominos
		allPlayersReady = 0
		for _, player := range current.GetPlayers() {
			if !player.HasStartingDominos() {
				chooseUnpickedDominos(current, player)
				fmt.Println(showBreak())
			} else {
				allPlayersReady++
			}
		}
	}
}

func chooseUnpickedDominos(current *game.Game, player *models.Player) bool {
	fmt.Print(showUnpickedDominos(current.GetTable()))
	fmt.Printf("Player %d\n", player.ID)
	return repeatUntilNoError(vaidatePickup, current, player.ID-1)
}

func chooseDominoToPlay(current *game.Game, player *models.Player) bool {
	fmt.Println(showBoard(current.GetBoard()))
	fmt.Println(showPlayerDominos(player))

	fmt.Printf("Player %d\n", player.ID)
	return repeatUntilNoError(validatePlay, current, player.ID-1)
}
