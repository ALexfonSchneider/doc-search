package entities

type Archive struct {
	ArchiveId string `json:"archive_id"`
	Name      string `json:"name"`
	Series    string `json:"series"`
	Url       string `json:"url"`
}
