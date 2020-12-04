package main

import (
	"awesomeProject/Handers"
	"awesomeProject/Job"
	"awesomeProject/MessageTypes"
	"net/http"
)

func main() {

	in := make(chan MessageTypes.Profile)

	out := make(chan MessageTypes.Profile)

	go Job.Workers(in, out)

	http.HandleFunc("/hello", Handers.Hello(in, out))
	http.ListenAndServe(":9000", nil)
}
