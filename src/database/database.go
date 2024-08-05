package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Inventory struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	UserId      uuid.UUID `json:"userId"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

type User struct {
	Id          uuid.UUID   `json:"id"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Username    string      `json:"username"`
	Items       []Inventory `json:"items"`
	CreatedTime time.Time   `json:"createdTime"`
	UpdatedTime time.Time   `json:"updatedTime"`
}

const (
	user    = "florenthaxhiu"
	dbname  = "inventoryManagement"
	sslMode = "disable"
)

func ConnectToDatabase() *sql.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", user, dbname, sslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetInventory(db *sql.DB, id uuid.UUID) (*sql.Rows, error) {
	row, err := db.Query("SELECT $1 FROM Inventory", id)

	if err != nil {
		return nil, errors.New("item not found")
	}

	return row, nil
}

func GetProfile(db *sql.DB, id uuid.UUID) User {
	var user User

	profile := db.QueryRow("SELECT $1 FROM 'Users'", id)
	profile.Scan(&user)

	fmt.Printf("%v+\n %v+\n", profile, user)

	return user
}

func PostInventory(db *sql.DB) {}
