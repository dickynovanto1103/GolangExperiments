package callback

import (
	"auth0/app"
	"auth0/auth"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("CALLBACK received")
	session, err := app.Store.Get(r, "auth-session")
	log.Printf("session: %v, err: %v", session, err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("url: %v r.URL.Get(state): %v session[state]: %v", r.URL.String(), r.URL.Query().Get("state"), session.Values["state"])

	if r.URL.Query().Get("state") != session.Values["state"] {
		http.Error(w, "invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("r.URL.Query(code): %v", r.URL.Query().Get("code"))

	token, err := authenticator.Config.Exchange(context.Background(), r.URL.Query().Get("code"))
	tokenBytes, _ := json.Marshal(token)

	log.Printf("token: %v, err: %v", string(tokenBytes), err)
	if err != nil {
		log.Printf("token not found, err: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	log.Printf("rawIDToken: %v", rawIDToken)
	if !ok {
		http.Error(w, "no id_token field in oauth2 token", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: "tCyyKW13VytGCsTaOTwBTm1HZ0lOuqH9",
	}

	log.Printf("verify rawIDToken")

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.Background(), rawIDToken)
	idTokenBytes, _ := json.Marshal(idToken)
	log.Printf("idToken: %v, err: %v", string(idTokenBytes), err)
	if err != nil {
		http.Error(w, "failed to verify idToken:"+err.Error(), http.StatusInternalServerError)
		return
	}

	//getting user info
	log.Printf("getting user info, get profile, idToken.Claims(&profile)")
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("profile: %v", profile)

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("session save successful, redirect to /user")

	//redirect to logged in page
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
