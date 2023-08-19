package post

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"get-cafedra.com/m/v2/types"
	"github.com/joho/godotenv"
)

func PostJson(web types.WebData, jsonBody []byte, api string) types.Response {
	client := web.Client
	token := web.Token
	url := web.Url //is subject to change
	path := "/api/" + api + "s/"
	req, err := http.NewRequest(http.MethodPost, url+path, bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Print("post.PostJson: ")
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Print("post.PostJson: ")
		fmt.Println(err)
	}
	fmt.Print("post.PostJson: ")
	text, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print("post.PostJson: ")
		fmt.Println(err)
	}
	var responseBody types.Response
	err = json.Unmarshal(text, &responseBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Status)
	defer res.Body.Close()
	return responseBody
}

func TestPost(web types.WebData) types.Response {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("post.TestPost: ")
		fmt.Println(err)
	}
	testJson := `
	{
		"Data": {
					"Title": "Hello",
					
					"Content":"lmao"
								  }
		}
	`
	return PostJson(web, []byte(testJson), "departament")
}

func GetHueten(web types.WebData) {
	client := web.Client
	token := web.Token
	url := web.Url //is subject to change
	path := "/api/test-huetens/1"
	req, err := http.NewRequest(http.MethodGet, url+path, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Status)
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bodyBytes))
}

func SendPreview(web types.WebData, filePath string, refId int, ref, field string) error {
	path := "/api/upload/"
	client := web.Client
	token := web.Token
	// New multipart writer.
	body := &bytes.Buffer{}
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	writer := multipart.NewWriter(body)
	//fw, err := writer.CreateFormFile("files", filePath) // old method (it doesn't change Content-Type header for this field)

	partHeader := textproto.MIMEHeader{}
	partHeader.Add("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "files", filepath.Base(filePath)))
	partHeader.Add("Content-Type", types.GetContentType(filePath))
	fw, err := writer.CreatePart(partHeader)

	if err != nil {
		return err
	}
	err = os.Chdir(filepath.Dir(filePath))
	if err != nil {
		return err
	}
	file, err := os.Open(filepath.Base(filePath))
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	fw, err = writer.CreateFormField("refId")
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(refId)))
	if err != nil {
		return err
	}
	fw, err = writer.CreateFormField("ref")
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(ref)))
	if err != nil {
		return err
	}
	fw, err = writer.CreateFormField("field")
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(field)))
	if err != nil {
		return err
	}
	writer.Close()
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, web.Url+path, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "bearer "+token)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)

	// respDump, err := httputil.DumpResponse(rsp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Printf("RESPONSE:\n%s", string(respDump))
	if rsp.StatusCode != http.StatusOK {
		text, err := io.ReadAll(rsp.Body)
		if err != nil {
			fmt.Print("post.SendFile: ")
			fmt.Println(err)
		}
		fmt.Println(string(text))
		return errors.New("Request failed with response code: " + fmt.Sprint(rsp.StatusCode))
	} else {
		return nil
	}
}

func SendMedia(web types.WebData, folderPath string, refId int, ref, field string) ([]types.ResponseMulti, error) {
	path := "/api/upload/"
	client := web.Client
	token := web.Token
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fs, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Chdir(folderPath)
	if err != nil {
		fmt.Println(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var fw io.Writer
	for i := 0; i < len(fs); i++ {
		partHeader := textproto.MIMEHeader{}
		partHeader.Add("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "files", filepath.Base(folderPath)+"_"+fs[i].Name()))
		partHeader.Add("Content-Type", types.GetContentType(folderPath))
		fw, err = writer.CreatePart(partHeader)

		if err != nil {
			return nil, err
		}
		file, err := os.Open(dir + `\` + fs[i].Name())
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			return nil, err
		}
	}
	fw, err = writer.CreateFormField("refId")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(refId)))
	if err != nil {
		return nil, err
	}
	fw, err = writer.CreateFormField("ref")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(ref)))
	if err != nil {
		return nil, err
	}
	fw, err = writer.CreateFormField("field")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(field)))
	if err != nil {
		return nil, err
	}
	writer.Close()
	if err := os.Chdir(currentDir); err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, web.Url+path, bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "bearer "+token)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)

	// respDump, err := httputil.DumpResponse(rsp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Printf("RESPONSE:\n%s", string(respDump))
	text, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Print("post.SendMedia: ")
		fmt.Println(err)
	}
	var responseBody []types.ResponseMulti
	err = json.Unmarshal(text, &responseBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Status)
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		text, err := io.ReadAll(rsp.Body)
		if err != nil {
			fmt.Print("post.SendFile: ")
			fmt.Println(err)
		}
		fmt.Println(string(text))
		return responseBody, errors.New("Request failed with response code: " + fmt.Sprint(rsp.StatusCode))
	} else {
		return responseBody, nil
	}
}

func UpdateArticleContent(web types.WebData, jsonBody []byte, api string, Id int) {
	client := web.Client
	token := web.Token
	url := web.Url //is subject to change
	path := "/api/" + api + "s/"
	req, err := http.NewRequest(http.MethodPut, url+path+fmt.Sprint(Id), bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Print("post.UpdateArticleContent: ")
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Print("post.UpdateArticleContent: ")
		fmt.Println(err)
	}
	fmt.Print("post.UpdateArticleContent: ")
	text, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print("post.UpdateArticleContent: ")
		fmt.Println(err)
	}
	var responseBody types.Response
	err = json.Unmarshal(text, &responseBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Status)
	defer res.Body.Close()

}
