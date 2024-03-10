package entities

type WordCloudItem struct {
	Value string `json:"value"`
	Count int32  `json:"count"`
}

type Metrics struct {
	WordCloud []WordCloudItem `json:"word_cloud"`
}

type Highlight struct {
	Content []string `json:"article.content"`
}

type SearchResult struct {
	Article   Article   `json:"article"`
	Archive   Archive   `json:"archive"`
	Metrics   Metrics   `json:"metrics"`
	Highlight Highlight `json:"highlight"`
}

type SearchResultsPaginate struct {
	Articles  []SearchResult `json:"articles"`
	Page      int            `json:"page"`
	Size      int            `json:"size"`
	TotalSize int            `json:"total_size"`
}
