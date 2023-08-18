package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"get-cafedra.com/m/v2/post"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	//depart := get.Departament(client)
	//get.ImageDepartament(client, depart)
	//teachers := get.Teachers(client)
	//get.ImageTeachers(client, teachers)
	//articles := get.Articles(client)
	//get.ImageArticles(client, articles)

	//DFull := get.DepartamentFull(client, depart)
	//TFull := get.TeachersFull(client, teachers)
	//AFull := get.ArticlesFull(client, articles)

	//fmt.Println(DFull[0].Content[0])
	//fmt.Println(TFull[0].Contacts[0])
	//get.ImageArticlesFull(client, AFull)

	resp := post.TestPost(client)
	post.GetHueten(client, os.Getenv("STRAPI_TOKEN2"))

	err := post.SendFile(client, "http://localhost:1337/api/upload/", "Downloaded/ArticlesPreview/60b21f5436fa3600a75a4d53.jpg", os.Getenv("STRAPI_TOKEN3"), resp.Data.Id, "api::departament.departament", "preview")
	if err != nil {
		fmt.Print("post.SendFile: ")
		fmt.Println(err)
	}
}
