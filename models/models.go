package models

type Deck map[string]int

type CardMap map[string]Card

type Card struct {
    Name string
    Type string
}
