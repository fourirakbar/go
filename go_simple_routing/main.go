package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Testing1"))
	}

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Again"))
	})

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/testing1", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = ":9000"
	fmt.Printf("server started at %s\n", address)

	server := new(http.Server)
	server.Addr = address

	// err := http.ListenAndServe(address, nil)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
