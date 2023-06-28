package types

import (
	"time"
)

type Depart struct {
	IsVisible bool   `json:"isVisible"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Preview   string `json:"preview"` //path to an image .jfif
}

type Article struct {
	IsVisible   bool
	CreatedAt   time.Time //response uses RFC 3339 or ISO-8601
	Id          string
	Title       string
	Preview     string //path to an image .jpg
	Description string //is always empty?
}

type ArticleResp struct {
	Articles []Article
	Total    int
}

type Teacher struct { //api sends []Teacher
	Id         string
	Avatar     string //path to an image
	Surname    string
	Firstname  string
	Patronymic string
	Position   string
}
