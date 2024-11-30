package main

import (
	"fmt"
	"net/http"

	"github.com/oziev02/url-shortener/configs"
	"github.com/oziev02/url-shortener/internal/auth"
	"github.com/oziev02/url-shortener/internal/link"
	"github.com/oziev02/url-shortener/internal/stat"
	"github.com/oziev02/url-shortener/internal/user"
	"github.com/oziev02/url-shortener/pkg/db"
	"github.com/oziev02/url-shortener/pkg/event"
	"github.com/oziev02/url-shortener/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
