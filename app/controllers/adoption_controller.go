package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dimasaldian/letsAdopt/app/config"
	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gorm.io/gorm"
)

func (server *Server) SubmitAdoption(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Ambil data dari form
	name := r.FormValue("name")
	email := r.FormValue("email")
	reason := r.FormValue("reason")
	petName := r.FormValue("petName") // Nama hewan

	// Cari ID hewan berdasarkan nama hewan
	var pet struct {
		ID uint
	}
	err = server.DB.Table("pets").Select("id").Where("name = ?", petName).First(&pet).Error
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}

	// Buat data adopsi
	adoption := models.Adoption{
		Name:   name,
		Email:  email,
		Reason: reason,
		PetID:  pet.ID, // ID hewan
		Status: "pending",
	}

	// Simpan data ke tabel Adoptions
	if err := server.DB.Create(&adoption).Error; err != nil {
		http.Error(w, "Failed to save adoption data", http.StatusInternalServerError)
		return
	}

	// Redirect ke halaman sukses
	http.Redirect(w, r, "/listpet", http.StatusSeeOther)
}

func (server *Server) AdoptionList(w http.ResponseWriter, r *http.Request) {
	var adoptions []models.Adoption

	err := server.DB.Preload("Pet").Find(&adoptions).Error
	if err != nil {
		http.Error(w, "Gagal untuk mengambil data adopsi", http.StatusInternalServerError)
		return
	}

	render := render.New(render.Options{
		Layout: "layout",
		Extensions: []string{".html"},
	})

	_ = render.HTML(w, http.StatusOK, "adoption_list", map[string]interface{}{
		"adoptions": adoptions,
		"activePage": "adoptions",
	})
}

func (server *Server) ApproveAdoption(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid adoption ID", http.StatusBadRequest)
        return
    }

    var adoption models.Adoption
    // Gunakan IDAdopt sebagai kolom primary key
    err = server.DB.Preload("Pet").First(&adoption, "id_adopt = ?", id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Adoption not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Failed to fetch adoption data", http.StatusInternalServerError)
        return
    }

    // Mulai transaksi untuk memastikan konsistensi
    tx := server.DB.Begin()

    // Perbarui status adopsi menjadi 'approved'
    if err := tx.Model(&adoption).Update("status", "approved").Error; err != nil {
        tx.Rollback()
        http.Error(w, "Gagal menyetujui adopsi", http.StatusInternalServerError)
        return
    }

    // Perbarui status hewan menjadi 'adopted'
    if err := tx.Model(&models.Pet{}).Where("id = ?", adoption.PetID).Update("status", "adopted").Error; err != nil {
        tx.Rollback()
        http.Error(w, "Gagal memperbarui status hewan", http.StatusInternalServerError)
        return
    }

    subject := "Pengajuan Adopsi Disetujui"
    body := fmt.Sprintf(`
        <h1>Selamat!</h1>
        <p>Pengajuan adopsi Anda untuk hewan <strong>%s</strong> telah disetujui.</p>
        <p>Silakan hubungi kami untuk langkah selanjutnya.</p>
    `, adoption.Pet.Name)

    if err := config.SendEmail(adoption.Email, subject, body); err != nil {
        tx.Rollback()
        http.Error(w, "Gagal mengirim email", http.StatusInternalServerError)
        return
    }


    tx.Commit()

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"Adopsi disetujui dan notifikasi terkirim"}`))
}



func (server *Server) RejectAdoption(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    
	var adoption models.Adoption
	err := server.DB.Preload("Pet").First(&adoption, "id_adopt = ?", id).Error
	if err != nil {
		http.Error(w, "Adoption not found", http.StatusNotFound)
		return
	}

	err = server.DB.Model(&adoption).Update("status", "rejected").Error
	if err != nil {
		http.Error(w, "Gagal menolak adopsi", http.StatusInternalServerError)
		return
	}
	subject := "Pengajuan Adopsi Ditolak"
	body := fmt.Sprintf(`
		<h1>Mohon Maaf</h1>
		<p>Pengajuan adopsi Anda untuk hewan <strong>%s</strong> telah ditolak.</p>
		<p>Jika Anda memiliki pertanyaan, silakan hubungi kami.</p>
	`, adoption.Pet.Name)

	err = config.SendEmail(adoption.Email, subject, body)
	if err != nil {
		http.Error(w, "gagal mengirim email", http.StatusInternalServerError)
		return
	}


    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Adopsi ditolak dan notifikasi terkirim"}`))
}