package types

import (
	"time"
)

type DepartDemo struct {
	IsVisible bool   `json:"isVisible"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Preview   string `json:"preview"` //path to an image .jfif
}

type DepartFull struct {
	IsVisible bool
	Id        string
	Title     string
	Preview   string //path to an image .jfif
	Content   string
	Media     []string //array of paths to the image
}

type ArticleDemo struct {
	IsVisible   bool
	CreatedAt   time.Time //response uses RFC 3339 or ISO-8601
	Id          string
	Title       string
	Preview     string //path to an image .jpg
	Description string //is always empty?
}

type ArticleFull struct {
	IsVisible   bool
	CreatedAt   time.Time //response uses RFC 3339 or ISO-8601
	Id          string
	Title       string
	Category    string
	Preview     string //path to an image .jpg
	Content     string
	Description string //is always empty?
	Author      string
	IsPublished bool
	UpdatedAt   time.Time
}

type ArticleResp struct {
	Articles []ArticleDemo
	Total    int
}

type TeacherDemo struct { //api sends []Teacher
	Id         string
	Avatar     string //path to an image
	Surname    string
	Firstname  string
	Patronymic string
	Position   string
}

type TeacherFull struct {
	Id          string
	Avatar      string //path to an image
	Surname     string
	Firstname   string
	Patronymic  string
	Position    string
	Education   string
	Courses     string
	Teaching    string
	Research    string
	Achivements string
	Info        string
	Contacts    string
	IsVisible   bool
}
