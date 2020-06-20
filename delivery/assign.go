package delivery

import (
	"fmt"

	"github.com/neiln3121/dominos/game"
)

// ChooseOption - Cycle through options for current game player
func ChooseOption(current *game.Game) error {
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
	fmt.Println("Success!")
	fmt.Println(showBreak())
	return nil
}

// GetPlayersReady - Get all players to pickup their initial dominos
func GetPlayersReady(current *game.Game) {
	var allPlayersReady int
	// Keep going until all players have the right amount of dominos
	for allPlayersReady < current.GetNumberOfPlayers() {
		// Set back to zero - every player needs the required amount of dominos
		allPlayersReady = 0
		for _, player := range current.GetPlayers() {
			if !player.HasStartingDominos() {
				chooseUnpickedDominos(current, player)
				fmt.Println(showPlayerDominos(player))
				fmt.Println(showBreak())
			} else {
				allPlayersReady++
			}
		}
	}
}
