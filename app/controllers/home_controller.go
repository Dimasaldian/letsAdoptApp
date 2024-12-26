package controllers

import (
	"net/http"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/unrolled/render"
)

// func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
// 	render := render.New(render.Options{
// 		Layout: "layout",
// 		Extensions: []string{".html"},
// 	})

// 	var pets []models.Pet
// 	result := server.DB.Debug().
// 		Preload("Images"). 
// 		Limit(4).    
// 		Find(&pets)

// 	if result.Error != nil {
// 		server.Logger.Printf("Error fetching pets: %v", result.Error)
// 		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
// 		return
// 	}
	
// 	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
// 		"title":      "Home Title",
// 		"body":       "Home Description",
// 		"showNavbar": true,
// 		"pets":       pets,
// 	})
// }
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
	})

	var pets []models.Pet

	// Query hanya untuk hewan yang tersedia
	result := server.DB.Debug().
		Preload("Images").
		Where("status = ?", "available"). // Filter hanya status 'available'
		Limit(4).                         // Batas jumlah data
		Find(&pets)

	if result.Error != nil {
		server.Logger.Printf("Error fetching pets: %v", result.Error)
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title":      "Home Title",
		"body":       "Home Description",
		"showNavbar": true,
		"pets":       pets,
		"showFooter": true,
	})
}

