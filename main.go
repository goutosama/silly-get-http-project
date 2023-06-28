package main

import (
	"net/http"

	get "get-cafedra.com/m/v2/get"
)

func main() {
	client := &http.Client{}
	depart := get.Departament(client)
	get.ImageDepartament(client, depart)
	teachers := get.Teachers(client)
	get.ImageTeachers(client, teachers)
	articles := get.Articles(client)
	get.ImageArticles(client, articles)
}
