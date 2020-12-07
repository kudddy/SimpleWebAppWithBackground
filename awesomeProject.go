package main

import "awesomeProject/Job"

//func main() {
//
//	r := mux.NewRouter()
//
//	// добавляем милдварю
//	r.Use(Handlers.JwtAuthentication)
//
//	in := make(chan MessageTypes.Profile)
//
//	out := make(chan MessageTypes.Profile)
//
//	go Job.Workers(in, out)
//
//	r.HandleFunc("/{token}/checktoken", Handlers.CheckToken)
//
//	r.HandleFunc("/status/{token}", Handlers.Hello(in, out))
//
//
//
//	http.Handle("/", r)
//	http.ListenAndServe(":9000", nil)
//}

func main() {
	token := "982be9592eb2dd614d58e341d323d96ceee54e1c"
	Job.AddFriendWorker(token)
}
