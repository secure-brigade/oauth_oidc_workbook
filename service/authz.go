package service

import (
	"context"
	"net/http"
)

type Authz interface {
	AuthorizeClient(w http.ResponseWriter, r *http.Request)
	IssueAccessToken(w http.ResponseWriter, r *http.Request)
}

type authz struct {
	ctx context.Context
}

func NewAuthzService(ctx context.Context) Authz {
	return &authz{ctx}
}

func (s *authz) AuthorizeClient(w http.ResponseWriter, r *http.Request)  {}
func (s *authz) IssueAccessToken(w http.ResponseWriter, r *http.Request) {}
