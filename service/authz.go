package service

import (
	"context"
	"net/http"
	"net/url"
)

type Authz interface {
	AuthorizeClient(w http.ResponseWriter, r *http.Request)
	IssueAccessToken(w http.ResponseWriter, r *http.Request)
}

type authz struct {
	ctx context.Context
}

type AuthorizeRequest struct {
}

func NewAuthzService(ctx context.Context) Authz {
	return &authz{ctx}
}

func (s *authz) AuthorizeClient(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	//clientID := params.Get("client_id")
	//responseType := params.Get("response_type")
	//scope := params.Get("scope")
	redirectUri, err := url.ParseRequestURI(params.Get("redirect_uri"))
	state := params.Get("state")
	//nonce := params.Get("nonce")
	//loginHint := params.Get("login_hint")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	q, err := url.ParseQuery(redirectUri.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	q.Add("code", "1234")
	q.Add("state", state)

	// https://tools.ietf.org/html/rfc6749#section-4.1.2
	w.WriteHeader(http.StatusFound)
	w.Header().Add("Location", q.Encode())
}

func (s *authz) IssueAccessToken(w http.ResponseWriter, r *http.Request) {}
