package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/oklog/ulid"
)

type MechModel struct {
	gorm.Model
	Uniq        string
	Name        string
	Version     string
	Release     string
	Description string
}

// NewULID returns a ULID as a string.
// NewULID func takes no as input and returns string
func NewULID() string {
	newid, _ := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	return newid.String()
}

// CreateDB creates db
func CreateDB() {
	db, err := gorm.Open("sqlite3", "mech.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&MechModel{})

	// Create
	id := NewULID()
	fmt.Printf("%s\n", id)
	db.Create(&MechModel{Uniq: id, Name: "foo", Version: "1.0.1", Release: "20191203", Description: "Mech of Brixton"})
	db.Create(&MechModel{Uniq: NewULID(), Name: "foo", Version: "1.0.0", Release: "20191203", Description: "Mech of Brixton"})

	// Read
	var first MechModel
	db.First(&first, "uniq = ?", id) // find product with id 1
	fmt.Printf("Name : %s\n", first.Name)
	fmt.Printf("Version : %s\n", first.Version)

	var sec MechModel
	db.First(&sec, "version = ?", "1.0.0")
	fmt.Printf("Name : %s\n", sec.Name)
	fmt.Printf("Version : %s\n", sec.Version)
	var mech MechModel
	db.First(&mech, "name = ?", "foo") // find product with code l1212
	fmt.Printf("Name : %s\n", mech.Name)
	fmt.Printf("Version : %s\n", mech.Version)
	// Update - update product's price to 2000
	db.Model(&mech).Update("Version", "1.0.2")

	var latest MechModel
	db.First(&latest, "version = ?", "1.0.2") // find product with id 1
	fmt.Printf("Name : %s\n", latest.Name)
	fmt.Printf("Version : %s\n", latest.Version)
	// Delete - delete product
	db.Delete(&latest)
}

func main() {
	CreateDB()
}
