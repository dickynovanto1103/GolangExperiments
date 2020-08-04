package home

import (
	"auth0/routes/templates"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
