package models

import (
	"github.com/akhidrb/toggl-cards/pkg/db"
	"github.com/lib/pq"
)

type Deck struct {
	db.Model
	Shuffled bool
	Cards    pq.StringArray `gorm:"type:varchar[]"`
}
