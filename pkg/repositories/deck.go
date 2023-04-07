package repositories

import (
	"errors"
	repositoriesI "github.com/akhidrb/toggl-cards/pkg/interfaces/repositories"
	"github.com/akhidrb/toggl-cards/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deck struct {
	db *gorm.DB
}

func NewDeck(dbConn *gorm.DB) repositoriesI.IDeck {
	return Deck{db: dbConn}
}

func (p Deck) Create(deck *models.Deck) error {
	return p.db.Create(deck).Error
}

func (p Deck) GetByID(id uuid.UUID) (*models.Deck, error) {
	deck := models.Deck{}
	err := p.db.First(&deck, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &deck, err
}

func (p Deck) Update(model models.Deck) error {
	return p.db.Model(&model).Updates(model).Error
}
