package services_test

import (
	"github.com/akhidrb/toggl-cards/pkg/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeck_ConstructCardList(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			expectedCards := []string{
				"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS", "AD",
				"2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD", "AC", "2C",
				"3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC", "AH", "2H", "3H",
				"4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
			}
			actualCards := services.ConstructCardList()
			assert.Equal(t, expectedCards, actualCards)
		},
	)
}

func TestDeck_ShuffleCards(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			cards := []string{
				"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS", "AD",
				"2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD", "AC", "2C",
				"3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC", "AH", "2H", "3H",
				"4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
			}
			shuffledCards := make([]string, len(cards))
			copy(shuffledCards, cards)
			services.ShuffleCards(shuffledCards)
			assert.NotEqual(t, cards, shuffledCards)
			assert.Equal(t, len(cards), len(shuffledCards))
		},
	)
}
