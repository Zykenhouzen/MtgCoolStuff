package card_game_engine

import "deck_statistics"

//  lists all possible states at the end of x turns and their probabilities
func GenerateStateProbabilities(deck deck_statistics.Deck, turns int) {
    //this function, in the future should be generated by a rule config for the
    //game. For now, it's just magic the gathering simplified rules and ai
    // (We just draw a card and pass)

    //We can totally do this recursively, but I want to try this out with
    //channels and threading, as that's the stuff this problem is made out of
}
