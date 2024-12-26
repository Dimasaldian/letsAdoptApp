package controllers

import (
	"log"
	"net/http"

	"github.com/Dimasaldian/letsAdopt/app/config"
	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Dashboard(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	adminName, _ := session.Values["adminName"].(string)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}

	renderer := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
	})

	adminModel := models.Admin{}
	admins, err := adminModel.GetAdmin(server.DB)
	if err != nil {
		log.Printf("Error retrieving admins: %v", err)
		http.Error(w, "Failed to load admins", http.StatusInternalServerError)
		return
	}

	err = renderer.HTML(w, http.StatusOK, "dashboard", map[string]interface{}{
		"title":      "Dashboard",
		"showNavbar": false,
		"adminName":    adminName,
		"admins":     admins,
	})

	if err != nil {
		log.Printf("Error rendering dashboard: %v", err)
		http.Error(w, "gagal merender dashboard", http.StatusInternalServerError)
		return
	}
}
