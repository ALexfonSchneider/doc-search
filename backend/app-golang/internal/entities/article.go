package entities

type Article struct {
	ArticleId string   `json:"article_id"`
	PreviewId string   `json:"preview_id"`
	Title     string   `json:"title"`
	Anatation string   `json:"anatation"`
	Keywords  []string `json:"keywords"`
	Udk       *string  `json:"udk,omitempty"`
	Published string   `json:"published"`
	Link      *string  `json:"link,omitempty"`
	Authors   []Author `json:"authors"`
}
