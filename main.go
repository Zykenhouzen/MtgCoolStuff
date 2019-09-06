package main

import (
	"fmt"
	"./models"
	"./deck_statistics"
)

func main() {
	// allCardsMap := models.CardMap{
	// 	"Forest" : models.Card{
	// 	Name: "Forest",
	// 	Type: "Land"}}

	startingDeck := models.Deck{
		"Forest": 30,
		"Island": 10,
	}
	fmt.Println("Started")
	chanceOfForest := deck_statistics.ChanceToDraw([]string{"Forest"}, startingDeck)
	fmt.Printf("Percent forest: %f%", (chanceOfForest*100))
}
