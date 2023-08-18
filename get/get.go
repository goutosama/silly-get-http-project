package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	types "get-cafedra.com/m/v2/types"
	"github.com/PuerkitoBio/goquery"
)

func Departament(web types.WebData) []types.DepartDemo {
	client := web.Client
	path := "/api/departament"
	resp, err := client.Get(web.UrlOld + path)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var departament []types.DepartDemo
		err = json.Unmarshal(bodyBytes, &departament)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return departament
	} else {
		fmt.Println(resp.StatusCode)
		return nil
	}
}

func Articles(web types.WebData) types.ArticleResp {
	client := web.Client
	path := "/api/article?limit=0&offset=0&orderBy=0"
	resp, err := client.Get(web.UrlOld + path)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var ArtResp types.ArticleResp
		err = json.Unmarshal(bodyBytes, &ArtResp)
		if err != nil {
			fmt.Println(err)
			return types.ArticleResp{}
		}
		return ArtResp
	} else {
		fmt.Println(resp.StatusCode)
		return types.ArticleResp{}
	}
}

func Teachers(web types.WebData) []types.TeacherDemo {
	client := web.Client
	path := "/api/teacher"
	resp, err := client.Get(web.UrlOld + path)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var teachers []types.TeacherDemo
		err = json.Unmarshal(bodyBytes, &teachers)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return teachers
	} else {
		fmt.Println(resp.StatusCode)
		return nil
	}
}

func ImageDepartament(web types.WebData, depart []types.DepartDemo) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("Downloaded"); err != nil {
		fmt.Println(err)
	}
	folderName := "Departament"
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir(folderName); err != nil {
		fmt.Println(err)
	}
	client := web.Client
	var url string = web.UrlOld
	for i := 0; i < len(depart); i++ {
		resp, err := client.Get(url + depart[i].Preview)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
		} else {
			fileName := depart[i].Id + "-" + depart[i].Title + ".jfif"
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}
}

func ImageArticles(web types.WebData, articles types.ArticleResp) { //needs rework for articles as whole
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("Downloaded"); err != nil {
		fmt.Println(err)
	}
	folderName := "ArticlesPreview"
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir(folderName); err != nil {
		fmt.Println(err)
	}
	client := web.Client
	var url string = web.UrlOld
	for i := 0; i < articles.Total; i++ {
		var resp *http.Response
		var err error
		if isExternalUrl(articles.Articles[i].Preview) {
			resp, err = client.Get(strings.ReplaceAll(articles.Articles[i].Preview, "\\", "/"))
		} else {
			resp, err = client.Get(url + strings.ReplaceAll(articles.Articles[i].Preview, "\\", "/"))
		}
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(fmt.Sprint(resp.StatusCode) + " " + articles.Articles[i].Id)
		} else {
			fileName := articles.Articles[i].Id + ".jpg"
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}
}

func ImageTeachers(web types.WebData, teachers []types.TeacherDemo) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("Downloaded"); err != nil {
		fmt.Println(err)
	}
	folderName := "Teachers"
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir(folderName); err != nil {
		fmt.Println(err)
	}
	client := web.Client
	var url string = web.UrlOld
	for i := 0; i < len(teachers); i++ {
		resp, err := client.Get(url + strings.ReplaceAll(teachers[i].Avatar, "\\", "/"))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
		} else {
			fileName := teachers[i].Id + "-" + teachers[i].Surname + "_" + teachers[i].Firstname + "_" + teachers[i].Patronymic + "_" + ".jpg"
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}
}

func isExternalUrl(url string) bool {
	return strings.HasPrefix(url, "http")
}

func DepartamentFull(web types.WebData, depart []types.DepartDemo) []types.DepartFull {
	client := web.Client
	var url string = web.UrlOld
	var path string = "/api/departament/"
	var result []types.DepartFull
	for i := 0; i < len(depart); i++ {
		resp, err := client.Get(url + path + depart[i].Id)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
		} else {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			var departament types.DepartFull
			err = json.Unmarshal(bodyBytes, &departament)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, departament)
		}
	}
	return result
}

func TeachersFull(web types.WebData, teacher []types.TeacherDemo) []types.TeacherFull {
	client := web.Client
	var url string = web.UrlOld
	var path string = "/api/teacher/"
	var result []types.TeacherFull
	for i := 0; i < len(teacher); i++ {
		resp, err := client.Get(url + path + teacher[i].Id)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
		} else {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			var teachers types.TeacherFull
			err = json.Unmarshal(bodyBytes, &teachers)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, teachers)
		}
	}
	return result
}

func ArticlesFull(web types.WebData, articles types.ArticleResp) []types.ArticleFull {
	client := web.Client
	var url string = web.UrlOld
	var path string = "/api/article/"
	var result []types.ArticleFull
	for i := 0; i < articles.Total; i++ {
		resp, err := client.Get(url + path + articles.Articles[i].Id)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
		} else {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			var article types.ArticleFull
			err = json.Unmarshal(bodyBytes, &article)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, article)
		}
	}
	return result
}

func ImageArticlesFull(web types.WebData, articles []types.ArticleFull) { //gets images inside Content field
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll("Downloaded", os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("Downloaded"); err != nil {
		fmt.Println(err)
	}
	folderName := "ArticlesFull"
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir(folderName); err != nil {
		fmt.Println(err)
	}
	//getting to the ArticlesFull folder
	client := web.Client
	var url string = web.UrlOld
	for i := 0; i < len(articles); i++ {
		fmt.Println(articles[i].Id)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(articles[i].Content))
		if err != nil {
			fmt.Println(err)
		}
		doc.Find("img").Each(func(c int, s *goquery.Selection) {
			text, b := s.Attr("src")
			if b {
				var resp *http.Response

				resp, err = client.Get(url + strings.Trim(text, "\\\""))

				if err != nil {
					fmt.Println(err)
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					fmt.Println(fmt.Sprint(resp.StatusCode) + " " + articles[i].Id)
				} else {
					currentDir, err := os.Getwd()
					if err != nil {
						fmt.Println(err)
					}
					folder := articles[i].Id
					if err := os.MkdirAll(folder, os.ModePerm); err != nil {
						fmt.Println(err)
					}
					if err := os.Chdir(folder); err != nil {
						fmt.Println(err)
					}
					fileName := fmt.Sprint(c) + ".jpg"
					file, err := os.Create(fileName)
					if err != nil {
						fmt.Println(err)
					}
					defer file.Close()
					_, err = io.Copy(file, resp.Body)
					if err != nil {
						fmt.Println(err)
					}
					if err := os.Chdir(currentDir); err != nil {
						fmt.Println(err)
					}
				}
			}
		})
	}
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}
}
