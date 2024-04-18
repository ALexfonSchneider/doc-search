package search

type Params struct {
	Query    string   `json:"q" query:"q"`
	Year     *int     `json:"year" query:"year"`
	Keywords []string `json:"keywords[]" query:"keywords[]"`
	Udk      *string  `json:"udk" query:"udk"`
	Page     int      `json:"page" query:"page" validate:"gt=0"`
	Size     int      `json:"count" query:"count" validate:"gt=0"`
}
