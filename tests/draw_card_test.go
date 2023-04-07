package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
)

func theUserDrawsCardsFromTheDeck(ctx context.Context, count int) (context.Context, error) {
	createdDeck, err := decodeCreatedDeckResponse(ctx)
	if err != nil {
		return ctx, err
	}
	url := fmt.Sprintf("deck/%s/draw", createdDeck.DeckID)
	params := map[string]string{"count": strconv.Itoa(count)}
	res, err := testConfig.adapter.Do(http.MethodPut, url, nil, params, nil)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, responseKey, res), err
}

func theUserShouldGetTheFollowingCards(ctx context.Context, cards *godog.Table) (
	context.Context, error,
) {
	res, ok := ctx.Value(responseKey).(*http.Response)
	if !ok {
		return ctx, errors.New("there is no response available")
	}
	if res.StatusCode >= 400 {
		err := errors.New("draw card from deck failed")
		return ctx, err
	}
	var actualResp DrawCardResponse
	err := json.NewDecoder(res.Body).Decode(&actualResp)
	if err != nil {
		return ctx, errors.New("failed to decode response")
	}
	for i, row := range cards.Rows {
		if i == 0 {
			continue
		}
		assert.Equal(&t, row.Cells[0].Value, actualResp.Cards[i-1].Code)
		if t.err != nil {
			return ctx, t.err
		}
		assert.Equal(&t, row.Cells[1].Value, actualResp.Cards[i-1].Value)
		if t.err != nil {
			return ctx, t.err
		}
		assert.Equal(&t, row.Cells[2].Value, actualResp.Cards[i-1].Suit)
		if t.err != nil {
			return ctx, t.err
		}
	}
	return ctx, nil
}
