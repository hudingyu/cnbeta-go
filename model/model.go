package model

type ArticleStruct struct {
	Sid          string `json:"sid" gorm:"primary_key;not null"`
	Title        string `json:"title"`
	PubTime      string `json:"inputtime"`
	Source       string `json:"source"`
	Url          string `json:"url_show"`
	Label        string `json:"label"`
	Hometext     string `json:"hometext"`
	Summary      string `json:"summary" gorm:"type:mediumText"`
	Content      string `json:"content" gorm:"type:longText"`
	Csrf         string `json:"csrf"`
	Sn           string `json:"sn"`
	Thumb        string `json:"thumb"`
	Author       string `json:"aid"`
	CommentCount string `json:"comments"`
}
