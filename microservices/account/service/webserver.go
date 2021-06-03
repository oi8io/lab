package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()    // NEW
	http.Handle("/", r) // NEW
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, r) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
