package repositories

import (
	repositoriesI "github.com/akhidrb/toggl-cards/pkg/interfaces/repositories"
	"github.com/akhidrb/toggl-cards/pkg/models"
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
