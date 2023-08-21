package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"get-cafedra.com/m/v2/transfer"
	"get-cafedra.com/m/v2/types"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	Web := types.WebData{
		Client: &http.Client{
			Timeout: time.Second * 120, //client will panic when uploading lots of files in one request in post.SendMedia
		},
		UrlOld: os.Getenv("OLD_URL"),
		Url:    os.Getenv("URL"),
		Token:  os.Getenv("STRAPI_TOKEN2"),
	}
	transfer.Departaments(Web)
	transfer.Teachers(Web)
	transfer.Articles(Web)

}
