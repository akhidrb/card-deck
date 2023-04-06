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
		dto.Cards = append(dto.Cards, constructCardDTOFromCode(card))
	}
}

type DrawCardsResponse struct {
	Cards []Card `json:"cards"`
}

func (dto *DrawCardsResponse) ModelToDTO(deck *models.Deck, count int) {
	*dto = DrawCardsResponse{}
	dto.Cards = make([]Card, 0, len(deck.Cards))
	for i, card := range deck.Cards {
		if i >= count {
			break
		}
		dto.Cards = append(dto.Cards, constructCardDTOFromCode(card))
	}
	deck.Cards = deck.Cards[count:]
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

func constructCardDTOFromCode(card string) Card {
	cardKeys := strings.Split(card, "")
	rankKey := cardKeys[0]
	suitKey := cardKeys[1]
	return Card{
		Value: cardRanksMap[rankKey],
		Suit:  cardSuitsMap[suitKey],
		Code:  card,
	}
}
