package sosalty

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	gifURL, err := scrape()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "<html><head></head><body><img src=\"%s\"></img></body></html>", gifURL)
}

func RunServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
