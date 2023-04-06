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
	"strconv"
	"strings"
)

const openDeck = "OPEN_DECK"

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
	return context.WithValue(ctx, openDeck, res), err
}

func theUserShouldReceiveADeckWithTheFollowingResults(
	ctx context.Context,
	results *godog.Table,
) (context.Context, error) {
	res, ok := ctx.Value(openDeck).(*http.Response)
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
	if t.err != nil {
		return ctx, t.err
	}
	expectedCards := strings.Split(results.Rows[1].Cells[1].Value, ",")
	if err != nil {
		return ctx, err
	}
	assert.Equal(&t, expectedCards, actualResp.Cards)
	return ctx, t.err
}
