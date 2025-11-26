package db

import (
	"database/sql"
	"fmt"
)

type Machine struct {
	Company      string `json:"Company"`
	MacAddress   string `json:"Mac_address"`
	SerialNumber string `json:"Serial_number"`
	IdModel      string `json:"id_model"`
}

type Collab struct {
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
}

type Model struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Brand    string `json:"brand"`
	Category string `json:"category"`
}

func CreateCollab(conn *sql.DB, collab *Collab) error {
	_, err := conn.Exec("INSERT INTO collab (first_name, last_name) VALUES ($1, $2)", collab.FirstName, collab.LastName)
	if err != nil {
		return err
	}
	return nil
}

func CreateMachine(conn *sql.DB, machine *Machine) error {
	_, err := conn.Exec("INSERT INTO machine (company, mac_address, serial_number, id_model) VALUES ($1, $2, $3, $4)", machine.Company, machine.MacAddress, machine.SerialNumber, machine.IdModel)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCollab(conn *sql.DB, id int) error {
	exists, err := CheckCollabExists(conn, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("aucun collaborateur trouvé avec l'id %d", id)
	}

	_, err = conn.Exec("DELETE FROM collab WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression du collaborateur : %v", err)
	}
	return nil
}

func DeleteMachine(conn *sql.DB, id int) error {
	exists, err := CheckMachineExists(conn, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("aucune machine trouvée avec l'id %d", id)
	}

	_, err = conn.Exec("DELETE FROM machine WHERE id = $1", id)
	return err
}

func CheckIdModel(conn *sql.DB, id int) (bool, error) {
	var exists bool
	err := conn.QueryRow("SELECT EXISTS(SELECT 1 FROM model WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckCollabExists(conn *sql.DB, id int) (bool, error) {
	var exists bool
	err := conn.QueryRow("SELECT EXISTS(SELECT 1 FROM collab WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func CheckMachineExists(conn *sql.DB, id int) (bool, error) {
	var exists bool
	err := conn.QueryRow("SELECT EXISTS(SELECT 1 FROM machine WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetAllCollabs(conn *sql.DB) ([]Collab, error) {
	rows, err := conn.Query("SELECT last_name,first_name FROM collab")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var collabs []Collab

	for rows.Next() {
		var c Collab
		if err := rows.Scan(&c.LastName, &c.FirstName); err != nil {
			return nil, err
		}
		collabs = append(collabs, c)
	}
	return collabs, nil
}

func GetAllMachine(conn *sql.DB) ([]Machine, error) {
	rows, err := conn.Query("SELECT company, mac_address, serial_number, id_model FROM machine")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var machines []Machine

	for rows.Next() {
		var m Machine
		if err := rows.Scan(&m.Company, &m.MacAddress, &m.SerialNumber); err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}

	return machines, nil
}

func GetCollab(conn *sql.DB, id int) (Collab, error) {
	var c Collab
	err := conn.QueryRow("SELECT last_name, first_name FROM collab WHERE id = $1", id).Scan(&c.LastName, &c.FirstName)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aucun collaborateur trouvé
			return Collab{}, nil
		}
		return Collab{}, err
	}
	return c, nil
}

func GetMachine(conn *sql.DB, id int) (Machine, error) {
	var m Machine
	err := conn.QueryRow("SELECT company, mac_address, serial_number, id_model FROM machine WHERE id = $1", id).Scan(&m.Company, &m.MacAddress, &m.SerialNumber, &m.IdModel)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aucun collaborateur trouvé
			return Machine{}, nil
		}
		return Machine{}, err
	}
	return m, nil
}

func UpdateCollab(conn *sql.DB, id int, d *Collab) error {
	exists, err := CheckCollabExists(conn, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("aucun collaborateur trouvé avec l'id %d", id)
	}

	_, err = conn.Exec("UPDATE collab SET first_name = $1, last_name = $2 WHERE id = $3", d.FirstName, d.LastName, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMachine(conn *sql.DB, id int, d *Machine) error {
	exists, err := CheckMachineExists(conn, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("aucune machine trouvée avec l'id %d", id)
	}

	_, err = conn.Exec("UPDATE machine SET company = $1, mac_address = $2, serial_number = $3, id_model = $4 WHERE id = $5",
		d.Company, d.MacAddress, d.SerialNumber, d.IdModel, id)
	if err != nil {
		return err
	}
	return nil
}
