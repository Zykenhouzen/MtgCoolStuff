package card_game_engine

import (
    "fmt"
    "../deck_statistics"
)

type MagicState struct {
    Turn int
    Hand deck_statistics.CardBag
    Deck deck_statistics.CardBag
}

type ProbAndState struct{
    Probability float64
    State MagicState
}

type MapOfStatesAndProbabilities map[string]*ProbAndState

//  lists all possible states at the end of x turns and their probabilities
func GenerateStateProbabilities(deck deck_statistics.CardBag, turns int) {
    //this function, in the future should be generated by a rule config for the
    //game. For now, it's just magic the gathering simplified rules and ai
    // (We just draw a card and pass)

    //We can totally do this recursively, but I want to try this out with
    //channels and threading, as that's the stuff this problem is made out of

}

func GenerateInitialStates(deck deck_statistics.CardBag) {

}

func DrawACard(state MagicState) (
    possibleStatesMap MapOfStatesAndProbabilities) {
        var drawProbabilities map[string]float64
        drawProbabilities = deck_statistics.ChanceToDrawEach(state.Deck)
        stateProbabilities := MapOfStatesAndProbabilities{}

        for cardName, cardProbability := range(drawProbabilities) {
            newDeck := deck_statistics.CardBag(state.Deck)
            newDeck[cardName] = newDeck[cardName] - 1

            newHand := deck_statistics.CardBag(state.Hand)
            newHand[cardName] = newHand[cardName] + 1

            newState := MagicState{
                Turn: state.Turn,
                Deck: newDeck,
                Hand: newHand,
            }
            // we need a hash of the whole state
            if _, exists:= stateProbabilities[newState.Deck.Hash()]; !exists{
                stateProbabilities[newState.Deck.Hash()] = &ProbAndState{
                    Probability: cardProbability,
                    State: newState,
                }
            } else {
                fmt.Println("Prob::: %f", stateProbabilities[newState.Deck.Hash()].Probability)
                stateProbabilities[newState.Deck.Hash()].Probability =
                    stateProbabilities[newState.Deck.Hash()].Probability +
                    cardProbability
            }
        }

        return stateProbabilities
}
