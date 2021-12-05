package main

import (
	"context"
	"log"
	"net/http"
	env "oauth-az/lib"
	"oauth-az/middlware"
	"oauth-az/service"

	"github.com/go-chi/chi"
)

func init() {
	env.SetupVars()
}

func main() {
	ctx := context.Background()

	router := chi.NewRouter()
	router.Use(middlware.CORS())
	router.Get("/hc", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	authorizeService := service.NewAuthzService(ctx)
	router.Get("/authorize", authorizeService.AuthorizeClient)
	router.Post("/token", authorizeService.IssueAccessToken)

	log.Println("connect to http" + "://" + env.GetURLAuthority())
	log.Fatal(http.ListenAndServe(env.GetURLAuthority(), router))
}
