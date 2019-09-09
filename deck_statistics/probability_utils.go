package deck_statistics

// send in a list of cards and a deck of cards
// returns the chance that you will draw one of the listed cards from the deck
func ChanceToDraw(cardNameList []string, deck Deck) float64 {
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

func ChanceToDrawEach(deck Deck) map[string]float64 {
	probabilityToDrawPerCard := map[string]float64{}
	var total int

	for _, amount := range deck {
		total += amount
	}

	for deckIndex, amount := range deck {
		 probabilityToDrawPerCard[deckIndex] = float64(amount)/float64(total)
	}

	return probabilityToDrawPerCard
}
