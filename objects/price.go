package objects

type Currency struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Price struct {
	Amount   int64    `json:"amount"`
	Currency Currency `json:"currency"`
	Text     string   `json:"text"`
}
