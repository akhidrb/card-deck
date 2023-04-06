package dtos

import (
	"github.com/akhidrb/toggl-cards/pkg/models"
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
}
