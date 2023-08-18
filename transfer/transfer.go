package transfer

import (
	"encoding/json"
	"fmt"
	"strings"

	"get-cafedra.com/m/v2/get"
	"get-cafedra.com/m/v2/post"
	types "get-cafedra.com/m/v2/types"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func HtmlToMarkdown(html string, MediaResponse []types.ResponseMulti) string {
	converter := md.NewConverter("", true, nil)
	if MediaResponse != nil {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(MediaResponse); i++ {
			doc.Find("img").Each(func(c int, s *goquery.Selection) {
				s.SetAttr("src", MediaResponse[c].Url)
				s.SetAttr("alt", MediaResponse[c].Name)
			})
		}
		html, err = doc.Html()
		if err != nil {
			fmt.Println(err)
		}
	}
	markdown, err := converter.ConvertString(html)
	if err != nil {
		fmt.Println(err)
	}
	return markdown
}

func Departaments(web types.WebData) {
	depart := get.Departament(web)
	get.ImageDepartament(web, depart)
	DFull := get.DepartamentFull(web, depart)
	get.ImageDepartamentFull(web, DFull)

	for i := 0; i < len(DFull); i++ {
		DFull[i].Content = HtmlToMarkdown(DFull[i].Content, nil)
		departNoFile := types.DepartNoFiles{
			IsVisible: DFull[i].IsVisible,
			Id:        DFull[i].Id,
			Title:     DFull[i].Title,
			Content:   DFull[i].Content,
		}
		var request types.Request = types.Request{
			Data: departNoFile,
		}
		jsonBody, err := json.Marshal(request)
		jsonBody[2] = byte('d')
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(string(jsonBody))
		resp := post.PostJson(web, jsonBody)
		err = post.SendPreview(web, `Downloaded/Departament/`+depart[i].Id+"-"+depart[i].Title+".jfif", resp.Data.Id, "api::departament.departament", "preview")
		if err != nil {
			fmt.Print("post.SendPreview: ")
			fmt.Println(err)
		}
		if len(DFull[i].Media) != 0 {
			_, err = post.SendMedia(web, `Downloaded/DepartFull/`+depart[i].Id, resp.Data.Id, "api::departament.departament", "media")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func Teachers(web types.WebData) {
	teachers := get.Teachers(web)
	get.ImageTeachers(web, teachers)
	TFull := get.TeachersFull(web, teachers)

	for i := 0; i < len(TFull); i++ {
		TFull[i].Education = HtmlToMarkdown(TFull[i].Education, nil)
		TFull[i].Courses = HtmlToMarkdown(TFull[i].Courses, nil)
		TFull[i].Teaching = HtmlToMarkdown(TFull[i].Teaching, nil)
		TFull[i].Research = HtmlToMarkdown(TFull[i].Research, nil)
		TFull[i].Achivements = HtmlToMarkdown(TFull[i].Achivements, nil)
		TFull[i].Info = HtmlToMarkdown(TFull[i].Info, nil)
		TFull[i].Contacts = HtmlToMarkdown(TFull[i].Contacts, nil)
		teacherNoFile := types.TeacherNoFiles{
			IsVisible:   TFull[i].IsVisible,
			Surname:     TFull[i].Surname,
			Firstname:   TFull[i].Surname,
			Patronymic:  TFull[i].Surname,
			Position:    TFull[i].Surname,
			Education:   TFull[i].Surname,
			Courses:     TFull[i].Surname,
			Teaching:    TFull[i].Surname,
			Research:    TFull[i].Surname,
			Achivements: TFull[i].Surname,
			Info:        TFull[i].Surname,
			Contacts:    TFull[i].Surname,
		}
		var request types.Request = types.Request{
			Data: teacherNoFile,
		}
		jsonBody, err := json.Marshal(request)
		jsonBody[2] = byte('d')
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(string(jsonBody))
		resp := post.PostJson(web, jsonBody)
		err = post.SendPreview(web, `Downloaded/Teachers/`+teachers[i].Id+"-"+teachers[i].Surname+"_"+teachers[i].Firstname+"_"+teachers[i].Patronymic+"_"+".jpg", resp.Data.Id, "api::teacher.teacher", "avatar")
		if err != nil {
			fmt.Print("post.SendPreview: ")
			fmt.Println(err)
		}
	}
}
