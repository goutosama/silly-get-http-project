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
	//testJson := `{"data":{"IsVisible":true,"CreatedAt":"0001-01-01T00:00:00Z","Id":"","Title":"Стань частью \"успешной\" команды!","Category":"news","Content":"**Давний партнер университета ООО «НПО Консультант» приглашает студентов кафедры Информационных технологий и цифровой экономики на вакантную должность администратора базы документов.**\n\nООО «НПО Консультант» – престижная, стабильная компания, активно ведущая свою деятельность больше 20 лет на региональном рынке в области информационных технологий. Ключевым направлением работы компании является распространение и поддержка справочной правовой системы СПС КонсультантПлюс.\n\n**Информация о вакансии:**\n\n**1.Должностные обязанности:**\n\n-Обработка текстовых документов согласно внутренним стандартам\n\n-Анализ взаимосвязей документов и создание редакций\n\n-Простановка гипертекстовых ссылок\n\n**2.График работы:** с понедельника по пятницу, с 8-00 до 17-00\n\n**3.Наличие испытательного срока:** испытательный срок 3 месяца\n\n**4.Заработная плата:** 26 000 – 28 000 рублей до вычета налогов\n\n**5.Социальный пакет:** официальное трудоустройство \\+ полное соблюдение ТК РФ: \"белая\" заработная плата, оплачиваемые отпуска и больничные\n\n**6.Другие условия работы:** своя столовая с льготными ценами для сотрудников, бесплатный ежегодный курс массажа, теннисный стол, яркая корпоративная жизнь.\n\n**Контактное лицо:**\n\nЛовцова Елена Михайловна, специалист по работе с персоналом\n\nТелефон 41-01-21\n\nE-mail [**lovczova@ivcons.ru**](mailto:lovczova@ivcons.ru)\n\n**Мы ценим профессионализм и высокую рабочую мотивацию. У нас психологически комфортно работать и развиваться, здесь чтут корпоративные традиции.**\n\n![6543c62171cc40dac763d3e32e.png](/uploads/2023/6/64a075860c4e7f99fbfc9cc8.png)","Description":"","Author":"6047470ee9ef1bb360d013d8","IsPublished":false,"UpdatedAt":"0001-01-01T00:00:00Z"}}`
	test2 := `{"data":{"IsVisible":true,"Title":"Ивановский филиал \"ЭнергосбыТ Плюс\" приглашает на работу молодых специалистов","Category":"news","Content":"**Крупнейшая энергосбытовая компания «ЭнергосбыТ Плюс» (Ивановский филиал) приглашает студентов кафедры Информационных технологий и цифровой экономики в свою команду!**\n\nАО «ЭнергосбыТ Плюс» – объединенная энергосбытовая компания Группы «Т Плюс» с филиальной сетью из 14 региональных филиалов на территории Российской Федерации. Компания объединяет энергию всей отрасли, создавая новые возможности для бизнеса и комфортные условия проживания для каждого из миллионов клиентов.\n\n![IMG_2827.jpg](/uploads/2023/5/6488e27fc80b5a834c82c719.jpg)\n\n**Большинство молодых специалистов уже сегодня задумываются о будущем. Все они амбициозны и рассчитывают на серьезный карьерный рост. Компания «ЭнергосбыТ Плюс» готова открыть для вас вполне реальные перспективы.**\n\n**На сегодняшний день в компании вакантны следующие должности:**\n\n**1\\. Специалист / Договорная группа / Управление балансов электрической энергии**\n\nДолжностные обязанности: заключение и сопровождение договоров оказания услуг по передаче электрической энергии; заключение и сопровождение договоров купли-продажи электрической энергии (мощности), приобретаемой сетевыми организациями в целях компенсации потерь.\n\nГрафик работы: 5/2 пн.-чт. с 8:00 до 17:00, пт. с 8:00 до 16:00\n\nИспытательный срок: 3 мес.\n\nЗаработная плата: 30 тыс.руб.\n\nСоц.пакет: есть\n\nДругие условия работы: возможность предоставления брони, ДМС\n\n**2\\. Специалист / Расчетно-экономическая группа / Управление балансов электрической энергии**\n\nДолжностные обязанности: расчет объема потерь электрической энергии, возникающих в электрических сетях сетевых организаций; формирование начислений по каждой сетевой организации; анализ и урегулирования разногласий в части объема потерь электрической энергии с сетевыми организациями; работа с дебиторской задолженностью, образованной в результате несвоевременного исполнения или неисполнения договорных обязательств (претензионно-исковая работа).\n\nГрафик работы: 5/2 пн.-чт. с 8:00 до 17:00, пт. с 8:00 до 16:00\n\nИспытательный срок: 3 мес.\n\nЗаработная плата: 35 тыс.руб.\n\nСоц.пакет: есть\n\nДругие условия работы: возможность предоставления брони, ДМС\n\n**Руководитель управления Положенцева Екатерина Валерьевна уже готова посвятить во все детали тех, кто заинтересовался. Телефон: +7 (4932) 93-73-78, E-mail:** [**Ekaterina.Polozhentseva@esplus.ru**](mailto:Ekaterina.Polozhentseva@esplus.ru)","Description":"","Author":"6047470ee9ef1bb360d013d8","IsPublished":false}}`

	return PostJson(web, []byte(test2), "article")
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
