package callback

import (
	"auth0/app"
	"auth0/auth"
	"context"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		http.Error(w, "invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		log.Printf("token not found, err: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "no id_token field in oauth2 token", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: "tCyyKW13VytGCsTaOTwBTm1HZ0lOuqH9",
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.Background(), rawIDToken)
	if err != nil {
		http.Error(w, "failed to verify idToken:"+err.Error(), http.StatusInternalServerError)
		return
	}

	//getting user info
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//redirect to logged in page
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
