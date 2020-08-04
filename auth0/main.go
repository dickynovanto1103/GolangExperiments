package main

import (
	"auth0/app"
	"auth0/routes/callback"
	"auth0/routes/home"
	"auth0/routes/login"
	"auth0/routes/logout"
	"auth0/routes/middlewares"
	"auth0/routes/user"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Init()
	if err != nil {
		log.Fatal(err)
	}
	server := mux.NewRouter()
	server.HandleFunc("/", home.HomeHandler)
	server.HandleFunc("/callback", callback.CallbackHandler)
	server.HandleFunc("/login", login.LoginHandler)
	server.HandleFunc("/logout", logout.LogoutHandler)

	server.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))

	server.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	log.Printf("listening to 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
