package sosalty

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The salt is real.")
}

func RunServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
