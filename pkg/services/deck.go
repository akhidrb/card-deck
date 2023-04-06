package services

import (
	"github.com/akhidrb/toggl-cards/pkg/dtos/requests"
	"github.com/akhidrb/toggl-cards/pkg/dtos/responses"
	repositoriesI "github.com/akhidrb/toggl-cards/pkg/interfaces/repositories"
	servicesI "github.com/akhidrb/toggl-cards/pkg/interfaces/services"
)

type Deck struct {
	repo repositoriesI.IDeck
}

func NewDeck(repo repositoriesI.IDeck) servicesI.IDeck {
	return Deck{repo: repo}
}

func (c Deck) Create(request requests.CreateDeckRequest) (res responses.CreateDeckResponse, err error) {
	return res, err
}
