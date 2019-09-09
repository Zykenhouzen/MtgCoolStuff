package deck_statistics

import (
    "testing"
    "reflect"
)

func TestChanceToDraw(t *testing.T) {
    tests := map[string]struct {
        cards []string
        deck  Deck
        want  float64
    }{
        "1 card out of 2":  {
            cards: []string{"Forest"},
            deck: Deck{
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

func TestChanceToDrawEach(t *testing.T) {
    tests := map[string]struct {
        cards []string
        deck  Deck
        want  map[string]float64
    }{
        "1 card out of 2":  {
            deck: Deck{
        		"Forest": 30,
        		"Island": 10,
        	},
            want: map[string]float64 {
                "Forest": .75,
                "Island": .25,
            },
        },
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            got := ChanceToDrawEach(tc.deck)
            if !reflect.DeepEqual(tc.want, got) {
                t.Fatalf("expected: %v, got: %v", tc.want, got)
            }
        })
    }
}
