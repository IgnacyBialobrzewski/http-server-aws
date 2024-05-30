package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("http server starting")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	log.Fatalln(http.ListenAndServe(":80", nil))
}
