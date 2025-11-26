package app

import (
	"INVENTORY_CGI/db"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

//struct de la table collab
//Column   |         Type          | Collation | Nullable |           Default
//------------+-----------------------+-----------+----------+------------------------------
//id         | integer               |           | not null | generated always as identity
//last_name  | character varying(50) |           | not null |
//first_name | character varying(50) |           | not null |

func GetAllCollabs(w http.ResponseWriter, r *http.Request) {
	collabs, err := db.GetAllCollabs(db.Conn)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(collabs)

}

func GetCollab(w http.ResponseWriter, r *http.Request) {
	//http.HandleFunc("GET /collabs/{id}", app.GetCollab)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	collabs, err := db.GetCollab(db.Conn, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(collabs)

}

func CreateCollab(w http.ResponseWriter, r *http.Request) {
	var collab db.Collab
	if err := json.NewDecoder(r.Body).Decode(&collab); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(collab.LastName, "<>[]{}*%") {
		http.Error(w, "Le nom à des characters interdit", http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(collab.FirstName, "<>[]{}*%") {
		http.Error(w, "Le nom à des characters interdit", http.StatusBadRequest)
		return
	}
	if err := db.CreateCollab(db.Conn, &collab); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func DeleteCollab(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := db.DeleteCollab(db.Conn, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateCollab(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var collab db.Collab
	if err := json.NewDecoder(r.Body).Decode(&collab); err != nil {
		http.Error(w, "Corps de requête invalide", http.StatusBadRequest)
		return
	}

	if strings.ContainsAny(collab.LastName, "<>[]{}*%") {
		http.Error(w, "Le nom contient des caractères interdits", http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(collab.FirstName, "<>[]{}*%") {
		http.Error(w, "Le prénom contient des caractères interdits", http.StatusBadRequest)
		return
	}

	if err := db.UpdateCollab(db.Conn, id, &collab); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Collaborateur mis à jour avec succès"))
}
