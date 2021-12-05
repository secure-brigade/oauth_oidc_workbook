package middlware

import (
	"net/http"
	env "oauth-az/lib"

	"github.com/rs/cors"
)

func CORS() func(h http.Handler) http.Handler {
	hoge := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead, http.MethodGet, http.MethodPost,
			http.MethodPut, http.MethodPatch, http.MethodDelete,
		},
		AllowedOrigins:   env.GetAllowedOrigins(),
		AllowCredentials: true,
		AllowedHeaders:   []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
	})
	return hoge.Handler
}
