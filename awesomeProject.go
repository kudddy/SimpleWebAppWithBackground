package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Profile struct {
	Name    string
	hobbies []string
}

func logLogs(logger chan Profile, data chan Profile) {
	print(logger)
	for item := range logger {
		fmt.Println("1. Item", item)
		time.Sleep(10*time.Second)
		item.Name = "Kirill"
		data <- item
	}
}

func hello(logger chan Profile, data chan Profile) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request){
		profile := Profile{"Alex", []string{"snowboarding", "programming"}}

		//logger <- profile

		select {
		case logger <- profile:
			fmt.Println("received message from hello", profile)
		default:
			fmt.Println("no message received from hello")
		}

		select {
		 case msg1 :=<-data:
		 	profile = msg1
		default:
			fmt.Println("io")
		}




		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	}
}


func main() {

	logs := make(chan Profile)

	data :=make(chan Profile)

	go logLogs(logs, data)

	handleHello := hello(logs, data)

	http.HandleFunc("/hello", handleHello)
	http.ListenAndServe(":9000", nil)
}

//func main() {
//	// METHOD 1
//	logs := make(chan string)
//	go logLogs(logs)
//	handleHello := makeHello(logs)
//
//	// METHOD 2
//	passer := &DataPasser{logs: make(chan string)}
//	go passer.log()
//
//	http.HandleFunc("/1", handleHello)
//	http.HandleFunc("/2", passer.handleHello)
//	http.ListenAndServe(":9999", nil)
//}
//
//// METHOD 1
//
//func makeHello(logger chan string) func(http.ResponseWriter, *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		logger <- r.Host
//		io.WriteString(w, "Hello world!")
//	}
//}
//
//func logLogs(logger chan string) {
//	for item := range logger {
//		fmt.Println("1. Item", item)
//	}
//}
//
//// METHOD 2
//
//type DataPasser struct {
//	logs chan string
//}
//
//func (p *DataPasser) handleHello(w http.ResponseWriter, r *http.Request) {
//	p.logs <- r.URL.String()
//	io.WriteString(w, "Hello world")
//}
//
//func (p *DataPasser) log() {
//	for item := range p.logs {
//		fmt.Println("2. Item", item)
//	}
//}

