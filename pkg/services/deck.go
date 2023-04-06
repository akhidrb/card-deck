package services

import (
	"fmt"
	"github.com/akhidrb/toggl-cards/pkg/dtos"
	repositoriesI "github.com/akhidrb/toggl-cards/pkg/interfaces/repositories"
	servicesI "github.com/akhidrb/toggl-cards/pkg/interfaces/services"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	repo repositoriesI.IDeck
}

func NewDeck(repo repositoriesI.IDeck) servicesI.IDeck {
	return Deck{repo: repo}
}

func (c Deck) Create(request dtos.CreateDeckRequest) (res dtos.CreateDeckResponse, err error) {
	if request.Cards != nil {
		request.CardsList = strings.Split(*request.Cards, ",")
	}
	if len(request.CardsList) == 0 {
		request.CardsList = c.constructCardList()
	}
	if request.Shuffle {
		c.shuffleCards(request.CardsList)
	}
	deck := request.ToModel()
	err = c.repo.Create(&deck)
	if err != nil {
		return dtos.CreateDeckResponse{}, err
	}
	res.ModelToDTO(deck)
	return
}

func (c Deck) GetByID(request dtos.OpenDeckRequest) (res dtos.OpenDeckResponse, err error) {
	id, _ := uuid.Parse(request.ID)
	deck, err := c.repo.GetByID(id)
	if err != nil {
		return res, err
	}
	res.ModelToDTO(deck)
	return
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
