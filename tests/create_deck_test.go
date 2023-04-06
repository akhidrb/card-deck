package tests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cucumber/godog"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"strings"
)

func aUserCreatesAFullDeckThatIsNotShuffled(ctx context.Context) (context.Context, error) {
	request := CreateDeckRequest{
		Shuffle: false,
		Cards:   nil,
	}
	return createDeck(ctx, request)
}

func aUserCreatesAFullDeckThatIsShuffled(ctx context.Context) (context.Context, error) {
	request := CreateDeckRequest{
		Shuffle: true,
		Cards:   nil,
	}
	return createDeck(ctx, request)
}

func aUserCreatesAPartialDeckThatIsNotShuffledWithTheFollowingCards(
	ctx context.Context,
	cardsTable *godog.Table,
) (context.Context, error) {
	cards := make([]string, 0, len(cardsTable.Rows))
	for _, row := range cardsTable.Rows {
		cards = append(cards, row.Cells[0].Value)
	}
	request := CreateDeckRequest{
		Shuffle: false,
		Cards:   cards,
	}
	return createDeck(ctx, request)
}

func createDeck(ctx context.Context, request CreateDeckRequest) (context.Context, error) {
	url := "deck"
	params := map[string]string{
		"shuffle": strconv.FormatBool(request.Shuffle),
	}
	if request.Cards != nil {
		params["cards"] = strings.Join(request.Cards, ",")
	}
	res, err := testConfig.adapter.Do(http.MethodPost, url, nil, params, nil)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, CreateDeckResponse{}, res), err
}

func theUserShouldReceiveADeckIDAndTheFollowingResults(
	ctx context.Context, results *godog.Table,
) (context.Context, error) {
	actualResp, err := decodeCreatedDeckResponse(ctx)
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
	return ctx, t.err
}

func decodeCreatedDeckResponse(ctx context.Context) (CreateDeckResponse, error) {
	var actualResp CreateDeckResponse
	res, ok := ctx.Value(CreateDeckResponse{}).(*http.Response)
	if !ok {
		return actualResp, errors.New("there is no response available")
	}
	if res.StatusCode >= 400 {
		err := errors.New("create deck failed")
		return actualResp, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	err := json.NewDecoder(res.Body).Decode(&actualResp)
	if err != nil {
		return actualResp, errors.New("failed to decode response")
	}
	return actualResp, err
}
