package main

import (
	"fmt"
	"github.com/oziev02/url-shortener/internal/hello"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
