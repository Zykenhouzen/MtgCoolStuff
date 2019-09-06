package deck_statistics

import (
	"../models"
)
// send in a list of cards and a deck of cards
// returns the chance that you will draw one of the listed cards from the deck
func ChanceToDraw(cardNameList []string, deck models.Deck) float64 {
	var totalFound, total int

	for _, name := range cardNameList {
		if(deck[name] != 0) {
			totalFound += deck[name]
		}
	}

	for _, amount := range deck {
		total += amount
	}

	return float64(totalFound)/float64(total)
}
