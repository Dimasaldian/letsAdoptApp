// package seeders

// import (
// 	"github.com/Dimasaldian/letsAdopt/database/fakers"
// 	"gorm.io/gorm"
// )

// type Seeder struct {
// 	Seeder interface{}
// }

// func RegisterSeeders(db *gorm.DB) []Seeder {
// 	return []Seeder{
// 		{Seeder: fakers.UserFaker(db)},
// 		{Seeder: fakers.PetFaker(db)},
// 	}
// }

// func DBSeed(db *gorm.DB) error {
// 	for _, seeder := range RegisterSeeders(db){
// 		err := db.Debug().Create(seeder.Seeder).Error
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

package seeders

import (
	"fmt"
	"log"

	"github.com/Dimasaldian/letsAdopt/database/fakers"
	"gorm.io/gorm"
)

func DBSeed(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		err := db.Debug().Create(fakers.UserFaker()).Error
		if err != nil {
			log.Fatal(err)
		}

		err = db.Debug().Create(fakers.PetFaker(db)).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database seeded successfully :)")
	return nil
}