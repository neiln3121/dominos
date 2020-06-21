package delivery

import (
	"fmt"
	"strings"

	"github.com/neiln3121/dominos/models"
)

func showUnpickedDominos(t *models.Table) string {
	var sb strings.Builder
	sb.WriteString("Unpicked dominos:\n")
	for i, domino := range t.GetDominos() {
		if !domino.IsPicked {
			sb.WriteString(fmt.Sprintf(" %d", i+1))
		}
	}
	sb.WriteString("\n\n")
	return sb.String()
}

func showPlayerDominos(p *models.Player) string {
	var sb strings.Builder
	sb.WriteString("Dominos: ")
	for i, domino := range p.GetDominos() {
		sb.WriteString(fmt.Sprintf("%d: %s ", i+1, renderDomino(domino)))
	}
	sb.WriteString("\n")
	return sb.String()
}

func renderDomino(d *models.Domino) string {
	half1, half2 := d.GetDots()
	return fmt.Sprintf("[%d|%d]", half1, half2)
}

func showBoard(b *models.Board) string {
	var sb strings.Builder
	sb.WriteString("Current board:\n-------------\n<-")
	for _, domino := range b.GetPlayedDominos() {
		sb.WriteString(fmt.Sprintf("%s-", renderDomino(domino)))
	}
	sb.WriteString(">\n\n")

	if b.GetHead() == b.GetTail() {
		sb.WriteString(fmt.Sprintf("You can only play: %d\n", b.GetHead()))
	} else {
		sb.WriteString(fmt.Sprintf("You can play either: %d or %d\n", b.GetHead(), b.GetTail()))
	}
	return sb.String()
}

func showBreak() string {
	return "---------------------------------\n"
}
