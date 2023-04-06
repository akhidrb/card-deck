package repositories

import "github.com/akhidrb/toggl-cards/pkg/models"

type IDeck interface {
	Create(model models.Deck) error
}
