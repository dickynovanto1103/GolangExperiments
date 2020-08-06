package logout

import (
	"log"
	"net/http"
	"net/url"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LOGOUT HANDLER")

	domain := "dickynovanto.au.auth0.com"
	logoutUrl, err := url.Parse("https://" + domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "/v2/logout"
	parameters := url.Values{}

	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}
	log.Printf("r.Host: %v remoteAddr: %v", r.Host, r.RemoteAddr)
	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("returnTo: %v", returnTo.String())
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", "tCyyKW13VytGCsTaOTwBTm1HZ0lOuqH9")
	logoutUrl.RawQuery = parameters.Encode()

	log.Printf("logoutUrl raw query: %v, redirect to logOutURL: %v", logoutUrl.RawQuery, logoutUrl.String())

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}
