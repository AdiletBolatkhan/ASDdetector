package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	http.HandleFunc("/", MainPage)
	http.Handle("/ui/html/", http.StripPrefix("/ui/html", http.FileServer(http.Dir("../../ui/html"))))
	http.Handle("/ui/css/", http.StripPrefix("/ui/css", http.FileServer(http.Dir("../../ui/css"))))
	fmt.Println("Running server on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println(err)
		return
	}
}
