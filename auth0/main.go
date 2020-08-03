package main

import (
	"auth0/routes/callback"

	"github.com/gorilla/mux"
)

func main() {
	server := mux.NewRouter()
	server.HandleFunc("/callback", callback.CallbackHandler)
}
