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
		_, err = post.SendMedia(web, `Downloaded/DepartFull/`+depart[i].Id, resp.Data.Id, "api::departament.departament", "media")
		if err != nil {
			fmt.Println(err)
		}
	}
}
