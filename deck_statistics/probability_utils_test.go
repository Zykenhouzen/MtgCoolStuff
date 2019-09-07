package deck_statistics

import (
    "testing"
    "reflect"
    "../models"
)

func TestChanceToDraw(t *testing.T) {
    tests := map[string]struct {
        cards []string
        deck  models.Deck
        want  float64
    }{
        "1 card out of 2":  {
            cards: []string{"Forest"},
            deck: models.Deck{
        		"Forest": 30,
        		"Island": 10,
        	},
            want: (.75),
        },
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            got := ChanceToDraw(tc.cards, tc.deck)
            if !reflect.DeepEqual(tc.want, got) {
                t.Fatalf("expected: %v, got: %v", tc.want, got)
            }
        })
    }
}