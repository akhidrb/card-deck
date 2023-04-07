package services

import (
	"errors"
	"fmt"
	"github.com/akhidrb/toggl-cards/pkg/config"
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
		if err = request.ValidateCards(); err != nil {
			return
		}
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
		return
	}
	res.ModelToDTO(deck)
	return
}

func (c Deck) GetByID(request dtos.OpenDeckRequest) (res dtos.OpenDeckResponse, err error) {
	id, _ := uuid.Parse(request.ID)
	deck, err := c.repo.GetByID(id)
	if err != nil {
		return
	}
	if deck == nil {
		err = config.NewNotFoundResourceError(errors.New("card deck does not exist"))
		return
	} else {
		res.ModelToDTO(*deck)
	}
	return
}

func (c Deck) DrawCards(request dtos.DrawCardsRequest) (res dtos.DrawCardsResponse, err error) {
	id, _ := uuid.Parse(request.ID)
	deck, err := c.repo.GetByID(id)
	if err != nil {
		return
	}
	res.ModelToDTO(deck, request.Count)
	if deck != nil {
		err = c.repo.Update(*deck)
		if err != nil {
			return
		}
	}
	return
}

var cardRanksList = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var cardSuitsList = []string{"S", "D", "C", "H"}

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
