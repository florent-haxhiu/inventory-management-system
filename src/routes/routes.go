package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/florent-haxhiu/inventory-management-system/v/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/google/uuid"
)

var db *sql.DB = database.ConnectToDatabase()

func Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Public Endpoints
	r.Group(func(r chi.Router) {
		r.Get("/inventory", getAllAssets)
		r.Get("/inventory/{id}", getAsset)

		r.Get("/profile/{id}", getProfile)

		r.Post("/login", Login)
		r.Post("/register", Register)
	})

	// Private Endpoints
	r.Group(func(r chi.Router) {})
	return r
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	var db = database.ConnectToDatabase()
	string_id := chi.URLParam(r, "id")
	id := uuid.MustParse(string_id)
	fmt.Printf("%s\n%s\n", string_id, id)
	profile := database.GetProfile(db, id)

	fmt.Fprintf(w, "Profile: %+v\n", profile)

	// w.Write([]byte(profile))
}

func getAllAssets(w http.ResponseWriter, r *http.Request) {

}

func getAsset(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Write([]byte(fmt.Sprintf("Id received %s\n", id)))
}

func createItem(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()

	w.Write([]byte(fmt.Sprintf("Id generated: %d", id)))
}
