package controllers

import (
	"errors"
	"net/http"

	"github.com/Dimasaldian/letsAdopt/app/config"
	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/unrolled/render"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminInput struct {
	Email    string
	Password string
}


func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout:     "layout",         // Menggunakan layout yang sama seperti Home
		Extensions: []string{".html"}, // File template dengan ekstensi .html
	})

	if r.Method == http.MethodGet {
		// Render halaman login
		err := render.HTML(w, http.StatusOK, "login", map[string]interface{}{
			"showNavbar": false,
		})
		if err != nil {
			server.Logger.Printf("Error rendering login page: %v", err)
			http.Error(w, "Gagal memuat halaman login", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		// Proses login
		r.ParseForm()
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		// Validasi input
		if email == "" || password == "" {
			_ = render.HTML(w, http.StatusBadRequest, "login", map[string]interface{}{
				"error":      "Email dan password wajib diisi.",
				"showNavbar": false,
				"showFooter": false,
			})
			return
		}

		// Cari admin berdasarkan email
		var admin models.Admin
		err := server.DB.Where("email = ?", email).First(&admin).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				_ = render.HTML(w, http.StatusUnauthorized, "login", map[string]interface{}{
					"error":      "Email atau password salah.",
					"showNavbar": false,
					"showFooter": false,
				})
				return
			}
			server.Logger.Printf("Error fetching admin: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Verifikasi password
		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
		if err != nil {
			_ = render.HTML(w, http.StatusUnauthorized, "login", map[string]interface{}{
				"error":      "Email atau password salah.",
				"showNavbar": false,
			})
			return
		}

		// Login berhasil, set session
		session, _ := config.Store.Get(r, config.SESSION_ID)
		session.Values["loggedIn"] = true
		session.Values["adminID"] = admin.ID
		session.Values["adminName"] = admin.Name
		session.Values["adminEmail"] = admin.Email
		session.Save(r, w)

		// Redirect ke dashboard
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	session.Values["loggedIn"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
