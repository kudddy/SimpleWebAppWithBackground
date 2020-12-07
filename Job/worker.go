package Job

import (
	"awesomeProject/MessageTypes"
	c "awesomeProject/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Workers(in chan MessageTypes.Profile, out chan MessageTypes.Profile) {
	for item := range in {
		fmt.Println("1. Item", item)
		time.Sleep(10 * time.Second)
		item.Name = "Kirill"
		out <- item
	}
}

// страндартный воркер должен иметь следующий функционал: выполнять джоб и писать в базу все действия
// Фронтальный компонент будет забирать статус из базы

func AddFriendWorker(token string) {
	// пытаемся распарсить профиль для получения основного языка програмирования
	// потом берем один из популярных фреймворком и получаем оттуда список юзеров к кому нужно добавиться
	// добавляем раз в несколько минут
	// Определяем язык клиента
	lang := "python"
	rand.Seed(time.Now().UnixNano())
	page := rand.Intn(700)

	fmt.Println(page)

	url := c.MapMainLangToRepo[lang] + "?page=" + strconv.Itoa(page)

	fmt.Println(url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	res, err := client.Do(req)
	fmt.Println(err)
	// TODOнужен еще один цил
	if res.StatusCode == 200 {
		//парсим json
		decoder := json.NewDecoder(res.Body)

		var resp []MessageTypes.StarGazers
		err := decoder.Decode(&resp)
		//fmt.Println(resp)
		if err == nil {
			for _, gazers := range resp {
				fmt.Println(gazers.Login)
				client := &http.Client{}
				fmt.Println(c.UrlAddFriend + gazers.Login)
				req, _ := http.NewRequest("PUT", c.UrlAddFriend+gazers.Login, nil)
				req.Header.Set("Accept", c.GitAccept)
				req.Header.Set("Authorization", "token"+" "+token)
				res, err := client.Do(req)
				if err == nil {
					// проверяем все ли успешно
					if res.StatusCode == 204 {
						fmt.Println("Пишем в базу что все успешно")
					} else {
						fmt.Println("Пишем в базу что не получилось")
					}
				} else {
					fmt.Println("Пишем в базу что сетевая ошибка")
				}
				//засыпаем
				time.Sleep(25 * time.Second)
			}
		}
	} else {
		fmt.Println("лалка")
	}
}
