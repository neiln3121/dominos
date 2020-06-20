package delivery

import (
	"fmt"

	"github.com/neiln3121/dominos/models"
)

func showUnpickedDominos(t *models.Table) string {
	str := "Unpicked dominos:\n"
	for i, domino := range t.GetDominos() {
		if !domino.IsPicked {
			str += fmt.Sprintf(" %d", i+1)
		}
	}
	return str + "\n"
}

func showPlayerDominos(p *models.Player) string {
	str := fmt.Sprintf("Player %d Dominos: ", p.ID)
	for i, domino := range p.GetDominos() {
		str += fmt.Sprintf("%d: %s ", i+1, renderDomino(domino))
	}
	return str + "\n"
}

func renderDomino(d *models.Domino) string {
	if d.PlayedFlipped {
		return fmt.Sprintf("[%d|%d]", d.Half[1], d.Half[0])
	}
	return fmt.Sprintf("[%d|%d]", d.Half[0], d.Half[1])
}

func showBoard(b *models.Board) string {
	str := "Current board:\n<-"
	for _, domino := range b.GetPlayedDominos() {
		str += fmt.Sprintf("%s-", renderDomino(domino))
	}
	str += fmt.Sprint(">\n\n")

	if b.GetHead() == b.GetTail() {
		str += fmt.Sprintf("You can only play: %d\n", b.GetHead())
	} else {
		str += fmt.Sprintf("You can play either: %d or %d\n", b.GetHead(), b.GetTail())
	}
	return str
}

func showBreak() string {
	return "---------------------------------\n"
}
