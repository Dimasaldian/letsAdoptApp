package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gorm.io/gorm"
)

func (server *Server) Listpet(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
	})

	var pets []models.Pet
	// Filter hanya hewan yang tersedia
	err := server.DB.Preload("Images").
		Where("status = ?", "available").
		Find(&pets).Error
	if err != nil {
		server.Logger.Printf("Error fetching pets: %v", err)
		http.Error(w, "Gagal mengambil data hewan", http.StatusInternalServerError)
		return
	}

	_ = render.HTML(w, http.StatusOK, "listpet", map[string]interface{}{
		"pets":       pets,
		"showNavbar": true,
		"showFooter": true,
	})
}


func (server *Server) AddPet(w http.ResponseWriter, r *http.Request) {
    render := render.New(render.Options{
        Layout: "layout",
        Extensions: []string{".html"},
    })

    _ = render.HTML(w, http.StatusOK, "add_pet", map[string]interface{}{
        "showNavbar": false,
    })
}

func (server *Server) CreatePet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Gagal memproses form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	petType := r.FormValue("type")
	breed := r.FormValue("breed")
	age, _ := strconv.Atoi(r.FormValue("age"))
	description := r.FormValue("description")
    negara := r.FormValue("negara")
	vaccinated := r.FormValue("vaccinated") == "true"

	images, err := uploadImages(r, "images")
	if err != nil {
		http.Error(w, "Gagal upload gambar", http.StatusInternalServerError)
		return
	}

	pet := models.Pet {
		Name:        name,
        Type:        petType,
        Breed:       breed,
        Age:         age,
        Description: description,
        Negara:     negara,
        Vaccinated:  vaccinated,
        Images:      images,
	}

	err = server.DB.Create(&pet).Error
	if err != nil {
		http.Error(w, "Gagal menyimpan data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard/listpet", http.StatusSeeOther)
}

func uploadImages(r *http.Request, fieldName string) ([]models.PetImage, error) {
	var petImages []models.PetImage

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	files := r.MultipartForm.File[fieldName]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		ext := filepath.Ext(fileHeader.Filename)
		newFilename := fmt.Sprintf("%s%s", generateRandomString(), ext)

		uploadDir := "assets/img/"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.MkdirAll(uploadDir, os.ModePerm)
		}
		filepath := filepath.Join(uploadDir, newFilename)

		dst, err := os.Create(filepath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return nil, err
		}

		petImage := models.PetImage{
			URL: "/public/img/" + newFilename,
		}
		petImages = append(petImages, petImage)
	}
	return petImages, nil
}

func generateRandomString() string {
	return uuid.New().String()
}

func (server *Server) PetDetailsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var pet models.Pet
	result := server.DB.Preload("Images").Where("name = ?", name).First(&pet)
	if result.Error != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}

	
	render := render.New(render.Options{
		Layout:      "layout",
		Extensions: []string{".html"},
	})

	_ = render.HTML(w, http.StatusOK, "pet_details", map[string]interface{}{
		"pet": pet,
		"showNavbar": true, 
		"showFooter": true,
	})
}

func (server *Server) EditPet(w http.ResponseWriter, r *http.Request) {
    render := render.New(render.Options{
        Layout:     "layout",
        Extensions: []string{".html"},
    })

    vars := mux.Vars(r)
    petID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid pet ID", http.StatusBadRequest)
        return
    }

    // Ambil data hewan berdasarkan ID
    var pet models.Pet
    err = server.DB.Preload("Images").First(&pet, petID).Error
    if err != nil {
        if gorm.ErrRecordNotFound == err {
            http.Error(w, "Pet not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Failed to fetch pet data", http.StatusInternalServerError)
        return
    }

    // Render halaman edit
    err = render.HTML(w, http.StatusOK, "edit_pet", map[string]interface{}{
        "pet":        pet,
        "showNavbar": false,
    })
    if err != nil {
        http.Error(w, "Failed to render edit page", http.StatusInternalServerError)
    }
}


func (server *Server) UpdatePet(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    petID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid pet ID", http.StatusBadRequest)
        return
    }

    // Ambil data lama dari database
    var pet models.Pet
    err = server.DB.Preload("Images").First(&pet, petID).Error
    if err != nil {
        if gorm.ErrRecordNotFound == err {
            http.Error(w, "Pet not found", http.StatusNotFound)
        } else {
            http.Error(w, "Error fetching pet data", http.StatusInternalServerError)
        }
        return
    }

    err = r.ParseMultipartForm(10 << 20) // Maksimum 10MB untuk gambar
    if err != nil {
        http.Error(w, "Failed to parse form data", http.StatusBadRequest)
        return
    }

    // Update nama
    if newName := r.FormValue("name"); newName != "" {
        pet.Name = newName
    }

    // Update jenis 
    if newType := r.FormValue("type"); newType != "" {
        validTypes := map[string]bool{
            "dog":  true,
            "cat":  true,
            "bird": true,
            "other": true,
        }
        if validTypes[newType] {
            pet.Type = newType
        } else {
            http.Error(w, "Invalid pet type", http.StatusBadRequest)
            return
        }
    }

    // Update ras 
    if newBreed := r.FormValue("breed"); newBreed != "" {
        pet.Breed = newBreed
    }

    // Update usia
    if newAge := r.FormValue("age"); newAge != "" {
        if age, err := strconv.Atoi(newAge); err == nil {
            pet.Age = age
        } else {
            http.Error(w, "Invalid age format", http.StatusBadRequest)
            return
        }
    }

    // Update deskripsi 
    if newDescription := r.FormValue("description"); newDescription != "" {
        pet.Description = newDescription
    }

    // Update status vaksinasi 
    if newVaccinated := r.FormValue("vaccinated"); newVaccinated != "" {
        pet.Vaccinated = newVaccinated == "true"
    }

    // Tambahkan gambar 
    if files, ok := r.MultipartForm.File["images"]; ok && len(files) > 0 {
        newImages, err := uploadImages(r, "images")
        if err != nil {
            http.Error(w, "Failed to upload images", http.StatusInternalServerError)
            return
        }
        pet.Images = append(pet.Images, newImages...)
    }

    // Simpan perubahan ke database
    err = server.DB.Save(&pet).Error
    if err != nil {
        server.Logger.Printf("Error updating pet: %v", err)
        http.Error(w, "Failed to update pet data", http.StatusInternalServerError)
        return
    }

    petCache.Delete("pets")

    http.Redirect(w, r, "/dashboard/listpet", http.StatusSeeOther)
}

func (server *Server) DeletePet(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    petID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid pet ID", http.StatusBadRequest)
        return
    }

    var pet models.Pet
    err = server.DB.Preload("Images").First(&pet, petID).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Pet not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Failed to fetch pet data", http.StatusInternalServerError)
        return
    }

    for _, image := range pet.Images {
        imagePath := "." + image.URL // Path gambar di server
        if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
            server.Logger.Printf("Failed to delete image file: %v", err)
        }
    }

    err = server.DB.Delete(&pet).Error
    if err != nil {
        server.Logger.Printf("Error deleting pet: %v", err)
        http.Error(w, "Failed to delete pet", http.StatusInternalServerError)
        return
    }

    petCache.Delete("pets")

    http.Redirect(w, r, "/dashboard/listpet", http.StatusSeeOther)
}

