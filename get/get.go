package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	types "get-cafedra.com/m/v2/types"
)

func Departament(client *http.Client) []types.Depart {
	resp, err := client.Get("https://oldit.isuct.ru/api/departament")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var departament []types.Depart
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

func Articles(client *http.Client) types.ArticleResp {
	resp, err := client.Get("https://oldit.isuct.ru/api/article?limit=0&offset=0&orderBy=0")
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

func Teachers(client *http.Client) []types.Teacher {
	resp, err := client.Get("https://oldit.isuct.ru/api/teacher")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var teachers []types.Teacher
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

func ImageDepartament(client *http.Client, depart []types.Depart) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("Downloaded"); err != nil {
		fmt.Println(err)
	}
	folderName := "Departament"
	if err := os.Mkdir(folderName, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir(folderName); err != nil {
		fmt.Println(err)
	}
	var url string = "https://oldit.isuct.ru"
	for i := 0; i < len(depart); i++ {
		resp, err := client.Get(url + depart[i].Preview)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
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
		} else {
			fmt.Println(resp.StatusCode)
		}
	}
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}
}

func ImageArticles(client *http.Client, articles types.ArticleResp) { //needs rework for articles as whole
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("Downloaded"); err != nil {
		fmt.Println(err)
	}
	folderName := "ArticlesPreview"
	if err := os.Mkdir(folderName, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir(folderName); err != nil {
		fmt.Println(err)
	}
	var url string = "https://oldit.isuct.ru"
	for i := 0; i < articles.Total; i++ {
		var resp *http.Response
		var err error
		if isExternalUrl(articles.Articles[i].Preview) {
			resp, err = client.Get(articles.Articles[i].Preview)
		} else {
			resp, err = client.Get(url + articles.Articles[i].Preview)
		}
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
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
		} else {
			fmt.Println(fmt.Sprint(resp.StatusCode) + " " + articles.Articles[i].Id)
		}
	}
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}
}

func ImageTeachers(client *http.Client, teachers []types.Teacher) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir("Downloaded"); err != nil {
		fmt.Println(err)
	}
	folderName := "Teachers"
	if err := os.Mkdir(folderName, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	if err := os.Chdir(folderName); err != nil {
		fmt.Println(err)
	}
	var url string = "https://oldit.isuct.ru"
	for i := 0; i < len(teachers); i++ {
		resp, err := client.Get(url + teachers[i].Avatar)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
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
		} else {
			fmt.Println(resp.StatusCode)
		}
	}
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}
}

func isExternalUrl(url string) bool {
	if string(url[0]) == "/" {
		return false
	} else {
		return true
	}
}
