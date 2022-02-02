package objects

type ButtonAction struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Button struct {
	Title  string       `json:"title"`
	Action ButtonAction `json:"action"`
}
