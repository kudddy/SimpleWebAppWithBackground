package handlers

import (
	"awesomeProject/MessageTypes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StarJobAdd(res http.ResponseWriter, req *http.Request) {
	// проверяем валидный ли токен
	vars := mux.Vars(req)

	token := vars["token"]

	fmt.Println(token)

	//tokenStatus:=realCheck(token)
	tokenStatus := true

	var workerStatus MessageTypes.CheckTokenResp

	workerStatus.MessageName = "STARTJOBADD"
	if tokenStatus {
		//запуск воркера
		fmt.Println("запуск воркера")
		workerStatus.Status = true
	} else {
		fmt.Println("Отправка сообщения о невалидности токена")
		workerStatus.Status = false
	}

	js, err := json.Marshal(workerStatus)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)

}
