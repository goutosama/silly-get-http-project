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
	//transfer.Departaments(Web)
	//transfer.Teachers(Web)
	//transfer.Articles(Web)

	//	transfer.HtmlToMarkdown(`{"isVisible":true,"createdAt":"2021-12-02T10:42:52.074Z","id":"61a8a32c7188ba2ff96f490c","title":"\tЗначимая победа в IT-хакатоне","category":"news","preview":"/uploads/2021/11/61a8a2627188ba28416f4906.jpg","content":"<p><strong>26-го и 27-го ноября в компании Akvelon прошел ежегодный хакатон для IT-разработчиков – Хаквелон 2021.&nbsp;В соревновании приняло участие 20 команд, состоящих из программистов-практиков, стажеров и студентов.</strong></p><p>В рамках 30 часового марафона по разработке программного обеспечения 60 участников работали над проектами, которые заключались в применении новых технологий в решении практико-ориентированных задач. Одним из условий участия в хакатоне было наличие хотя бы одного программиста от компании&nbsp;Akvelon&nbsp;в команде.</p><p><strong>В этом году в составе смешанных команд было много студентов из ИГХТУ. Это магистранты кафедры Информационных технологий и цифровой экономики:&nbsp;</strong>Егор Дунаев (2-130),&nbsp;Марат Зимнуров (2-130),&nbsp;Мария Карюгина (1-232) и&nbsp;Олег Войтович (1-233),&nbsp;а также&nbsp;<strong>студенты, обучающиеся по программам бакалавриата направления 09.03.02 Информационные системы и технологии:&nbsp;</strong>Семён Пиголицын (1-42),&nbsp;Александр Бутусов (1-147),&nbsp;Ярослав Вдовин (1-185),&nbsp;Игорь Амелин (2-147),&nbsp;Никита Рыбкин (2-147),&nbsp;Кирилл Шестаков (2-147),&nbsp;Руслан Арибжанов (2-147),&nbsp;Константин Тюкалов (4-42),&nbsp;Глеб Васягин (4-42),&nbsp;Данила Голубев (2-42),&nbsp;Александр Кириллов (2-42),&nbsp;Елизавета Ермакова (2-42),&nbsp;Никита Ужастин (2-184),&nbsp;Анна София Филиппова (3-42) и&nbsp;Юлия Струнникова (3-147).</p><p><strong>Студенты из ИГХТУ поучаствовали в разных проектах, таких как управление компьютером через жесты, сервис по предоставлению информации о проекте, проставление статусов во всех соц-сетях, cloud gaming, Project Matrix и ряде других.</strong></p><p>В результате защит проектов&nbsp;<strong>лучшим признан проект AR glasses, выполненный под руководством преподавателя-практика – разработчика компании&nbsp;Akvelon&nbsp;и доцента кафедры&nbsp;Информационных технологий и цифровой экономики&nbsp;Евгения Сергеевича Константинова. В составе победившей команды - 2 студента ИГХТУ Игорь Амелин и Егор Дунаев,</strong>&nbsp;который уже более года является также сотрудником кампании&nbsp;Akvelon.</p><p>Участник&nbsp;проект Project Matrix&nbsp;Никита Ужастин: \"Вот и закончился Хаквелон, хоть мы и не заняли призовое место, но 33 часа без сна прошли незабываемо: я получил просто колоссальное количество опыта во время работы с профессионалами, которые задали мне вектор развития и неубиваемую мотивацию двигаться дальше. По организации всё было прекрасно: еда, прекрасный кофе, рабочие места и оборудование. Хакатон в следующем году я посещу обязательно, буду побеждать!\"</p><p><strong>Поздравляем команду победителей и желаем дальнейших IT побед! А всем остальным студентам советуем развивать свой практический опыт, участвуя в различных мероприятиях!</strong></p><p><img src=\"/uploads/2021/11/61a8a2627188ba28416f4906.jpg\" alt=\"pVz1vB66mpg.jpg\"> <img src=\"/uploads/2021/11/61a8a2797188ba9e466f4907.jpg\" alt=\"fd7N-CFvVgw.jpg\"></p><p><img src=\"/uploads/2021/11/61a8a2967188ba10ae6f4908.jpg\" alt=\"изображение_viber_2021-12-01_23-50-50-540.jpg\"> <img src=\"/uploads/2021/11/61a8a2a37188ba34306f4909.jpg\" alt=\"awZI-KAxTUE.jpg\"></p><p><img src=\"/uploads/2021/11/61a8a2b07188baad9c6f490a.jpg\" alt=\"Ce8EwZJryj4.jpg\"> <img src=\"/uploads/2021/11/61a8a2b97188badd5e6f490b.jpg\" alt=\"YAK9XyWFEUQ.jpg\"></p>","description":"","author":"6047470ee9ef1bb360d013d8","isPublished":true,"updatedAt":"2021-12-02T10:42:52.074Z"}`)
	// articles := get.Articles(Web)
	// AFull := get.ArticlesFull(Web, articles)
	// var ASorted []types.ArticleFull
	// for i := 0; i < len(AFull); i++ {
	// 	if AFull[i].Id == "61a8a32c7188ba2ff96f490c" || AFull[i].Id == "6047498025860b38705c9879" {
	// 		ASorted = append(ASorted, AFull[i])
	// 	}
	// }
	//post.SendPreview(Web, `Downloaded/ArticlesPreview/`+"60da4cc83462767edfae0aa9"+".jpg", 2655, "api::article.article", "preview")
	transfer.Articles(Web)
}
