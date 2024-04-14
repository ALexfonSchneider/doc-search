package entity

type Archive struct {
	ArchiveId string `json:"archive_id" bson:"archive_id"`
	Name      string `json:"name" bson:"name"`
	Series    string `json:"series" bson:"series"`
	Url       string `json:"url" bson:"url"`
}
