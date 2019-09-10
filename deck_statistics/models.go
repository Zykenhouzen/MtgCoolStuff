package deck_statistics

type CardBag map[string]int

func (c CardBag) Hash() string {
    var cardHashSimple string
    for cardName, cardNum := range(c) {
        cardHashSimple = cardHashSimple + cardName + string(cardNum)
    }
    return cardHashSimple
}

type CardMap map[string]Card

type Card struct {
    Name string
    Type string
}
