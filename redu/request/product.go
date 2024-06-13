package request

type SearchProductListParam struct {
	Kw                string   `json:"kw"`
	Cm                string   `json:"cm"`
	Pp0               string   `json:"pp0"`
	Pp1               string   `json:"pp1"`
	F0                string   `json:"f0"`
	F1                string   `json:"f1"`
	Pf                string   `json:"pf"`
	By                string   `json:"by"`
	Sort              string   `json:"sort"`
	Platform          string   `json:"platform"`
	Type              string   `json:"type"`
	Page              int      `json:"page"`
	MemberId          string   `json:"member_id"`
	ProductCategoryId string   `json:"product_category_id"`
	ProductTagIds     []string `json:"product_tag_ids"`
	NotFake           bool     `json:"not_fake"`
}

var PageSize = 20
