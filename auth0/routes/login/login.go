package login

import (
	"auth0/app"
	"auth0/auth"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("LOGIN HANDLER")
	//generate random state
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := base64.StdEncoding.EncodeToString(bytes)
	log.Printf("state: %v", state)
	session, err := app.Store.Get(r, "auth-session")
	log.Printf("session: %v err: %v", session, err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("session save success")

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("from LOGIN HANDLER redirect to authcodeURL: %v", authenticator.Config.AuthCodeURL(state))
	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}
