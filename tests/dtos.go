package tests

type CreateDeckRequest struct {
	Shuffle bool     `json:"shuffle"`
	Cards   []string `json:"cards"`
}

type CreateDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}
