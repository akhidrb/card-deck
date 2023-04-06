package services

import (
	"github.com/akhidrb/toggl-cards/pkg/dtos"
)

type IDeck interface {
	Create(request dtos.CreateDeckRequest) (dtos.CreateDeckResponse, error)
}
