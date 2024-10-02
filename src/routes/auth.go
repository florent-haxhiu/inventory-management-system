package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/florent-haxhiu/inventory-management-system/v/database"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user database.User

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("User: %+v", user)

	fmt.Fprintf(w, "User: %+v", user)
}

func Login(w http.ResponseWriter, r *http.Request) {

}
