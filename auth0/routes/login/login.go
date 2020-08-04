package login

import (
	"auth0/app"
	"auth0/auth"
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//generate random state
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := base64.StdEncoding.EncodeToString(bytes)
	session, err := app.Store.Get(r, "auth-session")
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

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}
