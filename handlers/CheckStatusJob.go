package handlers

import (
	"awesomeProject/MessageTypes"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CheckStatusJob(res http.ResponseWriter, req *http.Request) {
	// генерация данных для проверки
	//tokenstatus:= false
	//status := MessageTypes.CheckTokenResp{MessageName: "TOKENSTATUS", Status: tokenstatus, StatusCode: 200}

	// достаем токен

	token := mux.Vars(req)["token"]

	ok := realCheck(token)
	var status MessageTypes.CheckJobStatusResp

	status.MessageName = "JOBSTATUS"

	// достаем статус задач

	result := models.GetJobStatusFromDb(token)

	print(result)

	if !ok {
		status.Status = result.Status

	} else {
		status.Status = result.Status
	}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)

}
