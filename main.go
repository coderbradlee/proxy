package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	port     string
	endpoint string
)

func get(uri string) ([]byte, int, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return ret, resp.StatusCode, nil
}

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request:", r.RequestURI)
	path := endpoint + r.RequestURI
	ret, statusCode, err := get(path)
	if err != nil {
		fmt.Println("response err:", statusCode, err)
		w.WriteHeader(statusCode)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(ret)
	fmt.Println("response:", statusCode, string(ret))
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
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
