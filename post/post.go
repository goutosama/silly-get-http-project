package post

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func PostJson(client *http.Client, token string, jsonBody []byte) {
	url := "http://localhost:1337" //is subject to change
	path := "/api/departaments"
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
	fmt.Println(res.Status + string(text))
}

func TestPost(client *http.Client) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("post.TestPost: ")
		fmt.Println(err)
	}
	testJson := `
	{
		"data": {
					"Title": "Hello",
					
					"Content":"lmao"
								  }
		}
	`
	PostJson(client, os.Getenv("STRAPI_TOKEN2"), []byte(testJson))
}

func GetHueten(client *http.Client, token string) {
	url := "http://localhost:1337" //is subject to change
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

func SendFile(urlPath, filePath, token string, refId int, ref, field string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("files", filePath)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
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

	req, err := http.NewRequest(http.MethodPost, urlPath, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		text, err := io.ReadAll(rsp.Body)
		if err != nil {
			fmt.Print("post.SendFile: ")
			fmt.Println(err)
		}
		fmt.Println(string(text))
		return errors.New("Request failed with response code:" + fmt.Sprint(rsp.StatusCode))
	} else {
		return nil
	}
}
