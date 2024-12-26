package controllers

import (
	"net/http"
	"sync"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/unrolled/render"
)

var petCache sync.Map

func (server *Server) DashboardListPet(w http.ResponseWriter, r *http.Request) {
    render := render.New(render.Options{
        Layout:     "layout",
        Extensions: []string{".html"},
    })

    var pets []models.Pet
    err := server.DB.Preload("Images").Find(&pets).Error
    if err != nil {
        server.Logger.Printf("Error fetching pets: %v", err)
        http.Error(w, "Gagal mengambil data hewan", http.StatusInternalServerError)
        return
    }

    _ = render.HTML(w, http.StatusOK, "dashboard_listpet", map[string]interface{}{
        "pets":       pets,
        "showNavbar": false,
    })
}