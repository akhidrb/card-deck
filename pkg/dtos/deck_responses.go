package dtos

import (
	"github.com/akhidrb/toggl-cards/pkg/models"
	"strings"
)

type CreateDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func (dto *CreateDeckResponse) ModelToDTO(deck models.Deck) {
	*dto = CreateDeckResponse{
		DeckID:    deck.ID.String(),
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
	}
}

type OpenDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func (dto *OpenDeckResponse) ModelToDTO(deck models.Deck) {
	*dto = OpenDeckResponse{
		DeckID:    deck.ID.String(),
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
	}
	dto.Cards = make([]Card, 0, len(deck.Cards))
	for _, card := range deck.Cards {
		cardKeys := strings.Split(card, "")
		rankKey := cardKeys[0]
		suitKey := cardKeys[1]
		dto.Cards = append(
			dto.Cards, Card{
				Value: cardRanksMap[rankKey],
				Suit:  cardSuitsMap[suitKey],
				Code:  card,
			},
		)
	}
}

var cardRanksMap = map[string]string{
	"A":  "ACE",
	"2":  "2",
	"3":  "3",
	"4":  "4",
	"5":  "5",
	"6":  "6",
	"7":  "7",
	"8":  "8",
	"9":  "9",
	"10": "10",
	"J":  "JACK",
	"Q":  "QUEEN",
	"K":  "KING",
}

var cardSuitsMap = map[string]string{
	"S": "SPADES",
	"D": "DIAMONDS",
	"C": "CLUBS",
	"H": "HEARTS",
}
