package yousosalty

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	gifURL, err := GetGifURL()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "<html><head></head><body><img height=\"100%%\" src=\"%s\"></img></body></html>", gifURL)
}

func RunServer() {
	log.Printf("Server Started at port 8080...")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
