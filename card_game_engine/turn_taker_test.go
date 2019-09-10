package card_game_engine

import (
    "testing"
    "reflect"
    "../deck_statistics"
)

func TestChanceToDraw(t *testing.T) {

    // we need to mock the deck_statisicts part here
    tests := map[string]struct {
        state  MagicState
        want  MapOfStatesAndProbabilities
    }{
        "Drawing 1 card from simple deck":  {
            state: MagicState{
                Turn: 0,
                Hand: deck_statistics.CardBag{},
                Deck: deck_statistics.CardBag{
                		"Forest": 30,
                		"Island": 10,
                    },
        	},
            want: nil,
        },
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            got := DrawACard(tc.state)
            if !reflect.DeepEqual(tc.want, got) {
                t.Fatalf("expected: %v, got: %v", tc.want, got)
            }
        })
    }
}
