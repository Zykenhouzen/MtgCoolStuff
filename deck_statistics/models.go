package deck_statistics

type CardBag map[string]int

type CardMap map[string]Card

type Card struct {
    Name string
    Type string
}
