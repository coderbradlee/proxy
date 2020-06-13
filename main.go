package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port     string
	endpoint string
)

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirect", endpoint)
	http.Redirect(w, r, endpoint, 301)
}
func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	endpoint = os.Getenv("ENDPOINT")
	if endpoint == "" {
		panic("endpoint should set through ENDPOINT env")
	}
	http.HandleFunc("*", redirect)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
