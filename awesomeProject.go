package main

import (
	"awesomeProject/Job"
	"awesomeProject/MessageTypes"
	"awesomeProject/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	// добавляем милдварю
	r.Use(handlers.JwtAuthentication)

	in := make(chan MessageTypes.Profile)

	out := make(chan MessageTypes.Profile)

	go Job.Workers(in, out)
	// скорее всего нужен будет менеджер

	r.HandleFunc("/{token}/checktoken", handlers.CheckToken)

	r.HandleFunc("/{token}/start/addfriend", handlers.StarJobAdd)

	r.HandleFunc("/status/{token}", handlers.Hello(in, out))

	http.Handle("/", r)
	http.ListenAndServe(":9000", nil)
}

//func main() {
//	token := "982be9592eb2dd614d58e341d323d96ceee54e1c"
//	Job.AddFriendWorker(token)
//}
