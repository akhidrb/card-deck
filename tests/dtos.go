package tests

type CreateDeckRequest struct {
	Shuffle bool
	Cards   []string
}

type CreateDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}
