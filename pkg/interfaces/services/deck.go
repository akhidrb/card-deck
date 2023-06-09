package services

import (
	"github.com/akhidrb/toggl-cards/pkg/dtos"
)

type IDeck interface {
	Create(request dtos.CreateDeckRequest) (dtos.CreateDeckResponse, error)
	GetByID(request dtos.OpenDeckRequest) (dtos.OpenDeckResponse, error)
	DrawCards(request dtos.DrawCardsRequest) (res dtos.DrawCardsResponse, err error)
}
