package main

import (
	"fmt"
	"net/http"

	"get-cafedra.com/m/v2/get"
)

func main() {
	client := &http.Client{}
	depart := get.Departament(client)
	get.ImageDepartament(client, depart)
	teachers := get.Teachers(client)
	get.ImageTeachers(client, teachers)
	articles := get.Articles(client)
	get.ImageArticles(client, articles)

	DFull := get.DepartamentFull(client, depart)
	TFull := get.TeachersFull(client, teachers)
	AFull := get.ArticlesFull(client, articles)

	fmt.Println(DFull[0].Content[0])
	fmt.Println(TFull[0].Contacts[0])
	get.ImageArticlesFull(client, AFull)

}
