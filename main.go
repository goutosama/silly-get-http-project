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
			Timeout: time.Second * 10,
		},
		UrlOld: os.Getenv("OLD_URL"),
		Url:    os.Getenv("URL"),
		Token:  os.Getenv("STRAPI_TOKEN2"),
	}
	transfer.Departaments(Web)
	// depart := get.Departament(Web)
	// DFull := get.DepartamentFull(Web, depart)
	// get.ImageDepartamentFull(Web, DFull)
	//teachers := get.Teachers(client)
	//get.ImageTeachers(client, teachers)
	//articles := get.Articles(client)
	//get.ImageArticles(client, articles)

	//TFull := get.TeachersFull(client, teachers)
	//AFull := get.ArticlesFull(client, articles)

	//fmt.Println(DFull[0].Content[0])
	//fmt.Println(TFull[0].Contacts[0])
	//get.ImageArticlesFull(client, AFull)

	//resp :=	post.TestPost(Web)
	// // post.GetHueten(Web)

	// err = post.SendPreview(Web, "Downloaded/ArticlesPreview/60b21f5436fa3600a75a4d53.jpg", resp.Data.Id, "api::departament.departament", "preview")
	// if err != nil {
	// 	fmt.Print("post.SendFile: ")
	// 	fmt.Println(err)
	// }

	// html := `
	// <p><strong>Всем известно, что профессии, связанные с управлением и информационными технологиями сегодня наиболее престижны и высокооплачиваемы. Но иногда ребятам непросто сделать выбор и определиться с перспективами построения карьеры в определенной сфере во время обучения в школе.</strong></p><p>За многолетний опыт кафедра информационных технологий и цифровой экономики нашла решение данной проблемы и всеми способами помогает абитуриентам в выборе будущей профессии.&nbsp;<strong>Для этого она активно ведет большую профориентационную работу со школьниками и студентами профессиональных образовательных организаций, интересующихся процессами применения цифровых технологий и менеджмента.</strong></p><p>За истекший год сотрудники кафедры посетили 3,9,67,41 и 65 школы города Иваново, Ивановский кооперативный техникум, Ивановский колледж сферы услуг, Ивановский автотранспортный колледж, Кинешемский технологический колледж, с рассказом о вузе и презентацией профилей и направлений подготовки бакалавров.</p><p>Помимо этого, на кафедре информационных технологий и цифровой экономики проводились интересные и увлекательные ИТ-мероприятия: Осенний марафон оСень++, мастер-классы по геймдеву и т.д. А на этой недели запланировано проведение олимпиады по информатике среди школьников и конференции учителей.</p><p>В свою очередь, мы не забываем и про просветительскую работу.&nbsp;<strong>Деятельность волонтеров цифрового просвещения, которую организует кафедра Информационных технологий и цифровой экономики, дает участникам возможность убедиться, что информатика и информационные технологии – это доступно для понимания, увлекательно, актуально и важно для нас всех.</strong>&nbsp;Ведь умение применять инновации делает нашу жизнь более легкой и интересной, а понимание рисков, угроз и способов их избежания делает цифровую трансформацию безопасной.</p><p><strong>Мы благодарим руководство колледжей и школ за теплый прием, а также сотрудников кафедры за проделанную профориентационную работу. В свою очередь, желаем всем школьникам успехов в подготовке к экзаменам и высоких баллов ЕГЭ! Мы будем искренне рады видеть всех, кто посетил наши&nbsp;IT-мероприятия, в стенах нашего вуза в рядах первокурсников 2021 года!</strong></p><p><img src=\"/uploads/2021/5/60da4e1b3462761235ae0aaa.jpg\" alt=\"x8rPc_0ryIA.jpg\"><img src=\"/uploads/2021/5/60da4e243462769f5dae0aab.jpg\" alt=\"Vh5YaYQ_yWM.jpg\"><img src=\"/uploads/2021/5/60da4e323462767c10ae0aac.jpg\" alt=\"_YGytwfw1js.jpg\"><img src=\"/uploads/2021/5/60da4e343462765b4aae0aad.jpg\" alt=\"GF0AR1VksR0.jpg\"></p>
	// `

	// resp2, err := post.SendMedia(Web, `Downloaded\ArticlesFull\60da4e3b3462765c34ae0aae`, 26, "api::departament.departament", "media")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(resp2[0].Url)

	// fmt.Println(transfer.HtmlToMarkdown(html, resp2))

	// folderPath := `Downloaded\ArticlesFull\60da4e3b3462765c34ae0aae`
	// fs, err := os.ReadDir(folderPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = os.Chdir(folderPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// dir, _ := os.Getwd()
	// file, err := os.Open(dir + `\` + fs[0].Name())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(file.Name())

}
