package repositories

import (
	"github.com/akhidrb/toggl-cards/pkg/models"
	"github.com/google/uuid"
)

type IDeck interface {
	Create(model *models.Deck) error
	GetByID(id uuid.UUID) (models.Deck, error)
	Update(model models.Deck) error
}
