package objects

type Product struct {
	Price Price `json:"price"`
}

type Link struct {
	Url         string  `json:"url"`
	Title       string  `json:"title"`
	Caption     string  `json:"caption"`
	Description string  `json:"description"`
	Photo       Photo   `json:"photo,omitempty"`
	Product     Product `json:"product,omitempty"`
	Button      Button  `json:"button,omitempty"`
	PreviewPage string  `json:"preview_page"`
	PreviewUrl  string  `json:"preview_url"`
}
