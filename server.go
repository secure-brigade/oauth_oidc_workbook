package main

import (
	"context"
	"log"
	"net/http"
	"oauth-az/lib/env"
	"oauth-az/lib/storage"
	"oauth-az/middlware"
	"oauth-az/service"

	"github.com/go-chi/chi"
)

func init() {
	env.SetupVars()
}

func main() {
	ctx := context.Background()

	c := storage.NewSQLite()
	defer c.Close()

	router := chi.NewRouter()
	router.Use(middlware.CORS())
	router.Get("/hc", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	authorizeService := service.NewAuthzService(ctx)
	router.Get("/authorize", authorizeService.AuthorizeRequest)
	router.Post("/token", authorizeService.TokenRequest)

	log.Println("connect to http" + "://" + env.GetURLAuthority())
	log.Fatal(http.ListenAndServe(env.GetURLAuthority(), router))
}
