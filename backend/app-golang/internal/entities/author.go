package entities

type Author struct {
	Affiliation string `json:"affiliation"`
	Name        string `json:"name"`
	Orcid       string `json:"orcid"`
}
