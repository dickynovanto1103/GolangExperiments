package auth

import (
	"context"
	"log"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

const (
	Domain = "https://dickynovanto.au.auth0.com/"
)

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, Domain)
	if err != nil {
		log.Printf("failed to get provider, err: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     "tCyyKW13VytGCsTaOTwBTm1HZ0lOuqH9",
		ClientSecret: "CLIENT SECRET", //TODO: use env variable
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:3000/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}
	authenticator := &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}
	return authenticator, nil
}
