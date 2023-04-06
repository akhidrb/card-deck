package tests

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

var testConfig TestConfig

type TestConfig struct {
	adapter Client
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		TestSuiteInitializer: InitTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func aUserCreatesAFullDeckThatIsNotShuffled(ctx context.Context) (context.Context, error) {
	url := "deck"
	request := CreateDeckRequest{
		Shuffled: false,
		Cards:    nil,
	}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return ctx, err
	}
	res, err := testConfig.adapter.Do(http.MethodPost, url, nil, nil, requestBody)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, CreateDeckResponse{}, res), err
}

func theUserShouldReceiveADeckIDAndTheFollowingResults(ctx context.Context, results *godog.Table) (context.Context, error) {
	res, ok := ctx.Value(CreateDeckResponse{}).(*http.Response)
	if !ok {
		return ctx, errors.New("there is no response available")
	}
	if res.StatusCode >= 400 {
		err := errors.New("create deck failed")
		return ctx, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	var actualResp CreateDeckResponse
	err := json.NewDecoder(res.Body).Decode(&actualResp)
	if err != nil {
		return ctx, err
	}
	expectedShuffled, err := strconv.ParseBool(results.Rows[1].Cells[0].Value)
	if err != nil {
		return ctx, err
	}
	assert.Equal(&t, expectedShuffled, actualResp.Shuffled)
	expectedRemaining, err := strconv.ParseInt(results.Rows[1].Cells[1].Value, 10, 0)
	if err != nil {
		return ctx, err
	}
	assert.Equal(&t, expectedRemaining, actualResp.Remaining)
	return ctx, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a user creates a full deck that is not shuffled$`, aUserCreatesAFullDeckThatIsNotShuffled)
	ctx.Step(`^the user should receive a deck ID and the following results:$`, theUserShouldReceiveADeckIDAndTheFollowingResults)
}

func InitTestSuite(ctx *godog.TestSuiteContext) {
	clientInst := NewClient("localhost:8080", time.Second*10)
	testConfig = TestConfig{adapter: clientInst}
}