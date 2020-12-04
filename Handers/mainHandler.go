package Handers

import (
	"awesomeProject/MessageTypes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Hello(in chan MessageTypes.Profile, out chan MessageTypes.Profile) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// Пример декодирования json
		decoder := json.NewDecoder(req.Body)
		var userinfo MessageTypes.UserToken
		err := decoder.Decode(&userinfo)
		if err != nil {
			panic(err)
		}
		fmt.Println(userinfo.Token)

		profile := MessageTypes.Profile{Name: "Alex", Hobbies: []string{"snowboarding", "programming"}}

		select {
		case in <- profile:
			fmt.Println("received message from hello", profile)
		case msg1 := <-out:
			profile = msg1
		default:
			fmt.Println("no message received from hello")
		}

		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	}
}