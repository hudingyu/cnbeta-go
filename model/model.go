package model

type ArticleStruct struct {
	Sid          string `json:"sid" gorm:"primary_key;not null"`
	Title        string `json:"title"`
	PubTime      string `json:"inputtime"`
	Source       string `json:"source"`
	Url          string `json:"url_show"`
	Label        string `json:"label"`
	Hometext     string
	Summary      string `gorm:"type:mediumText"`
	Content      string `gorm:"type:longText"`
	Csrf         string
	Sn           string
	Thumb        string `json:"thumb"`
	Author       string `json:"aid"`
	CommentCount string `json:"comments"`
}
