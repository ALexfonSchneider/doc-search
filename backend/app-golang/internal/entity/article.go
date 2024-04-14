package entity

type Article struct {
	ArticleId string   `json:"article_id" bson:"article_id"`
	PreviewId string   `json:"preview_id" bson:"preview_id"`
	Title     string   `json:"title" bson:"title"`
	Anatation string   `json:"anatation" bson:"anatation"`
	Keywords  []string `json:"keywords" bson:"keywords"`
	Udk       *string  `json:"udk,omitempty" bson:"udk"`
	Published string   `json:"published" bson:"published"`
	Content   *string  `json:"content,omitempty" bson:"content"`
	Link      *string  `json:"link,omitempty" bson:"link"`
	Authors   []Author `json:"authors" bson:"authors"`
}
