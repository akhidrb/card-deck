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

func theUserShouldOpenADeckWithFollowingResults(
	ctx context.Context, results *godog.Table,
) (context.Context, error) {
	var actualResp OpenDeckResponse
	res, ok := ctx.Value(responseKey).(*http.Response)
	if !ok {
		return ctx, errors.New("there is no response available")
	}
	if res.StatusCode >= 400 {
		err := errors.New("create deck failed")
		return ctx, err
	}
	err := json.NewDecoder(res.Body).Decode(&actualResp)
	if err != nil {
		return ctx, errors.New("failed to decode response")
	}
	if err != nil {
		return ctx, err
	}
	expectedShuffled, err := strconv.ParseBool(results.Rows[1].Cells[0].Value)
	if err != nil {
		return ctx, err
	}
	assert.Equal(&t, expectedShuffled, actualResp.Shuffled)
	if t.err != nil {
		return ctx, t.err
	}
	expectedRemaining, err := strconv.Atoi(results.Rows[1].Cells[1].Value)
	if err != nil {
		return ctx, err
	}
	assert.Equal(&t, expectedRemaining, actualResp.Remaining)
	return context.WithValue(ctx, responseKey, actualResp), t.err
}

func theCardsInTheDeckShouldBe(ctx context.Context, cards *godog.Table) (
	context.Context,
	error,
) {
	actualResp, ok := ctx.Value(responseKey).(OpenDeckResponse)
	if !ok {
		return ctx, errors.New("there is no response available")
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
