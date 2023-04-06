package services

import (
	"fmt"
	"github.com/akhidrb/toggl-cards/pkg/dtos"
	repositoriesI "github.com/akhidrb/toggl-cards/pkg/interfaces/repositories"
	servicesI "github.com/akhidrb/toggl-cards/pkg/interfaces/services"
	"math/rand"
	"time"
)

type Deck struct {
	repo repositoriesI.IDeck
}

func NewDeck(repo repositoriesI.IDeck) servicesI.IDeck {
	return Deck{repo: repo}
}

func (c Deck) Create(request dtos.CreateDeckRequest) (res dtos.CreateDeckResponse, err error) {
	if len(request.Cards) == 0 {
		request.Cards = c.constructCardList()
	}
	if request.Shuffle {
		c.shuffleCards(request.Cards)
	}
	deck := request.ToModel()
	err = c.repo.Create(&deck)
	if err != nil {
		return dtos.CreateDeckResponse{}, err
	}
	res.ModelToDTO(deck)
	return res, err
}

func (c Deck) constructCardList() []string {
	cardList := make([]string, 0, 52)
	for _, suit := range cardSuitsList {
		for _, rank := range cardRanksList {
			value := fmt.Sprintf("%s%s", rank, suit)
			cardList = append(cardList, value)
		}
	}
	return cardList
}

func (c Deck) shuffleCards(cards []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}
