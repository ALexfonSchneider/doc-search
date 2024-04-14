package entity

type WordCloudItem struct {
	Value string `json:"value" bson:"value"`
	Count int32  `json:"count" bson:"count"`
}

type DocumentsInYearCount struct {
	Year  int32 `bson:"year" json:"year"`
	Count int32 `bson:"count" json:"count"`
}

type Metrics struct {
	WordCloud []WordCloudItem        `json:"word_cloud" bson:"word_cloud"`
	Years     []DocumentsInYearCount `json:"years,omitempty" bson:"-"`
}
