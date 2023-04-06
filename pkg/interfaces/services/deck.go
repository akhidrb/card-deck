package services

import (
	"github.com/akhidrb/toggl-cards/pkg/dtos/requests"
	"github.com/akhidrb/toggl-cards/pkg/dtos/responses"
)

type IDeck interface {
	Create(request requests.CreateDeckRequest) (responses.CreateDeckResponse, error)
}
