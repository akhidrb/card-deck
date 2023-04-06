package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func theUserRequestsToOpenTheCreatedDeck(ctx context.Context) (context.Context, error) {
	createdDeck, err := decodeCreatedDeckResponse(ctx)
	if err != nil {
		return ctx, err
	}
	url := fmt.Sprintf("deck/%s", createdDeck.DeckID)
	res, err := testConfig.adapter.Do(http.MethodGet, url, nil, nil, nil)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, responseKey, res), err
}

func theCardsInTheDeckShouldBe(ctx context.Context, cards *godog.Table) (
	context.Context,
	error,
) {
	res, ok := ctx.Value(responseKey).(*http.Response)
	if !ok {
		return ctx, errors.New("there is no response available")
	}
	if res.StatusCode >= 400 {
		err := errors.New("open deck failed")
		return ctx, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	var actualResp OpenDeckResponse
	err := json.NewDecoder(res.Body).Decode(&actualResp)
	if err != nil {
		return ctx, errors.New("failed to decode response")
	}
	for i, row := range cards.Rows {
		if i == 0 {
			continue
		}
		assert.Equal(&t, row.Cells[0].Value, actualResp.Cards[i])
		if t.err != nil {
			return ctx, t.err
		}
	}
	return ctx, nil
}
