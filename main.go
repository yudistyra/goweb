package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(resp http.ResponseWriter, req *http.Request) {
	var message = "Welcome"
	resp.Write([]byte(message))
}

func handlerHello(resp http.ResponseWriter, req *http.Request) {
	var message = "Hello World"
	resp.Write([]byte(message))
}

func main() {
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	// http.HandleFunc("/hello", handlerHello)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	var address = "localhost:9000"
	fmt.Printf("Server started at %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
