// package fakers

// import (
// 	"math/rand"
// 	"time"

// 	"github.com/Dimasaldian/letsAdopt/app/models"
// 	"github.com/bxcodec/faker/v3"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// func PetFaker(db *gorm.DB) *models.Pet {
// 	rand.Seed(time.Now().UnixNano())

// 	petTypes := []string{"dog", "cat", "bird", "other"}
// 	petType := petTypes[rand.Intn(len(petTypes))]

// 	// Jika tipe hewan peliharaan adalah 'other', kosongkan breed
// 	var breed string
// 	if petType != "other" {
// 		breed = faker.Word()
// 	}

// 	return &models.Pet{
// 		ID:          uuid.New().String(),
// 		Name:        faker.FirstName(),
// 		Type:        petType,
// 		Breed:       breed,
// 		Age:         rand.Intn(15), // Umur antara 0-14 tahun
// 		Description: faker.Sentence(),
// 		Vaccinated:  rand.Intn(2)==0,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}
// }
package fakers

import (
	"math/rand"
	"time"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func PetFaker(db *gorm.DB) *models.Pet {
	rand.Seed(time.Now().UnixNano())

	petTypes := []string{"dog", "cat", "bird", "other"}
	petType := petTypes[rand.Intn(len(petTypes))]

	// Jika tipe hewan peliharaan adalah 'other', kosongkan breed
	var breed string
	if petType != "other" {
		breed = faker.Word()
	}

	return &models.Pet{
		Name:        faker.FirstName(),
		Type:        petType,
		Breed:       breed,
		Age:         rand.Intn(15), // Umur antara 0-14 tahun
		Description: faker.Sentence(),
		Vaccinated:  rand.Intn(2) == 0,
	}
}