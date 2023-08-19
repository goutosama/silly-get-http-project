package types

import (
	"net/http"
	"path/filepath"
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

type Response struct {
	Data Data
}

type Data struct {
	Id         int
	Attributes map[string]string
}

type Request struct {
	Data interface{}
}

type ResponseMulti struct {
	Id              int
	Name            string
	AlternativeText string
	Caption         string
	Width           int
	Height          int

	Hash              string
	Ext               string
	Mime              string
	Size              float64
	Url               string
	ReviewUrl         string
	Provider          string
	Provider_metadata string //?
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
type DataMulti struct {
	Data string
}

func GetContentType(filePath string) string {
	ext := filepath.Ext(filePath)
	var res string
	if ext == ".jpg" || ext == ".jpeg" || ext == ".jpe" || ext == ".jfif" {
		res = "jpeg"
	} else {
		switch ext {
		case ".gif":
			res = "gif"
		case ".png":
			res = "png"
		case ".tiff":
			res = "tiff"
		case ".svg":
			res = "svg+xml"
		case ".svgz":
			res = "svg+xml"
		case ".ico":
			res = "vnd.microsoft.icon"
		case ".wbmp":
			res = "vnd.wap.wbmp"
		case ".webp":
			res = "webp"
		}
	}

	return "image/" + res
}

type WebData struct {
	Client *http.Client
	UrlOld string
	Url    string
	Token  string
}

type DepartNoFiles struct {
	IsVisible bool
	Id        string
	Title     string
	Content   string
}

type TeacherNoFiles struct {
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

type ArticleNoFiles struct {
	IsVisible bool
	CreatedAt time.Time //response uses RFC 3339 or ISO-8601
	Id        string
	Title     string
	Category  string

	Content     string
	Description string //is always empty?
	Author      string
	IsPublished bool
	UpdatedAt   time.Time
}
