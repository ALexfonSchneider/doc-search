package entity

type Highlight struct {
	Content []string `json:"article.content"`
}

type Document struct {
	Article   Article    `json:"article" bson:"article"`
	Archive   Archive    `json:"archive" bson:"archive"`
	Metrics   Metrics    `json:"metrics" bson:"metrics"`
	Highlight *Highlight `json:"highlight,omitempty" bson:"-"`
	Deleted   bool       `json:"-" bson:"deleted"`
}

type SearchResultsPaginate struct {
	Articles  []Document `json:"articles"`
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	TotalSize int        `json:"total_size"`
}
