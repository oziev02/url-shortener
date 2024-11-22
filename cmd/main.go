package main

import (
	"fmt"
	"net/http"

	"github.com/oziev02/url-shortener/configs"
	"github.com/oziev02/url-shortener/internal/auth"
	"github.com/oziev02/url-shortener/internal/link"
	"github.com/oziev02/url-shortener/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

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
