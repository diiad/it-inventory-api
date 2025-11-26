package app

//                                   Table "public.machine"
//    Column     |         Type          | Collation | Nullable |           Default
//---------------+-----------------------+-----------+----------+------------------------------
// id            | integer               |           | not null | generated always as identity
// company       | character varying(50) |           | not null |
// mac_address   | character varying(30) |           |          | NULL::character varying
// serial_number | character varying(30) |           | not null |
// id_model      | integer               |           | not null |
//---------------+-----------------------+-----------+----------+------------------------------
//Indexes:
//    "machine_pkey" PRIMARY KEY, btree (id)
//Foreign-key constraints:
//    "machine_id_model_fkey" FOREIGN KEY (id_model) REFERENCES model(id)
//Referenced by:
//    TABLE "soclage_event" CONSTRAINT "soclage_event_id_machine_fkey" FOREIGN KEY (id_machine) REFERENCES machine(id)

import (
	"INVENTORY_CGI/db"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var DB *sql.DB

func GetAllMachine(w http.ResponseWriter, r *http.Request) {
	machine, err := db.GetAllMachine(db.Conn)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(machine)
}

func GetMachine(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	machines, err := db.GetMachine(db.Conn, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(machines)

}

func CreateMachine(w http.ResponseWriter, r *http.Request) {
	var machine db.Machine
	if err := json.NewDecoder(r.Body).Decode(&machine); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(machine.Company, "<>[]{}*%") {
		http.Error(w, "Le nom à des characters interdit", http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(machine.MacAddress, "<>[]{}*%") {
		http.Error(w, "Le nom à des characters interdit", http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(machine.SerialNumber, "<>[]{}*%") {
		http.Error(w, "Le nom à des characters interdit", http.StatusBadRequest)
		return
	}
	if _, err := strconv.Atoi(machine.IdModel); err != nil {
		http.Error(w, "Le champ id_model doit être un entier", http.StatusBadRequest)
		return
	}

	var idModel int
	idModel, _ = strconv.Atoi(machine.IdModel)

	exists, err := db.CheckIdModel(db.Conn, idModel)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !exists {
		http.Error(w, "L'id_model n'existe pas", http.StatusBadRequest)
		return
	}

	if err := db.CreateMachine(db.Conn, &machine); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func UpdateMachine(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var machine db.Machine
	if err := json.NewDecoder(r.Body).Decode(&machine); err != nil {
		http.Error(w, "Corps de requête invalide", http.StatusBadRequest)
		return
	}

	if strings.ContainsAny(machine.Company, "<>[]{}*%") {
		http.Error(w, "Le nom de l'entreprise contient des caractères interdits", http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(machine.MacAddress, "<>[]{}*%") {
		http.Error(w, "L'adresse MAC contient des caractères interdits", http.StatusBadRequest)
		return
	}
	if strings.ContainsAny(machine.SerialNumber, "<>[]{}*%") {
		http.Error(w, "Le numéro de série contient des caractères interdits", http.StatusBadRequest)
		return
	}

	idModel, err := strconv.Atoi(machine.IdModel)
	if err != nil {
		http.Error(w, "Le champ id_model doit être un entier", http.StatusBadRequest)
		return
	}

	exists, err := db.CheckIdModel(db.Conn, idModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "L'id_model n'existe pas", http.StatusBadRequest)
		return
	}

	if err := db.UpdateMachine(db.Conn, id, &machine); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Machine mise à jour avec succès"))
}

func DeleteMachine(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := db.DeleteMachine(db.Conn, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
