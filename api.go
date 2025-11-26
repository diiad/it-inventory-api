package main

import (
	"INVENTORY_CGI/app"
	"INVENTORY_CGI/db"
	"fmt"
	"net/http"
)

func main() {
	db.Conn = db.NewDB()
	http.HandleFunc("GET /{$}", healthCheck)

	http.HandleFunc("POST /machine", app.CreateMachine)
	http.HandleFunc("DELETE /machine/{id}", app.DeleteMachine)
	http.HandleFunc("GET /machine", app.GetAllMachine)
	http.HandleFunc("GET /machine/{id}", app.GetMachine)
	http.HandleFunc("PATCH /machine/{id}", app.UpdateMachine)

	http.HandleFunc("POST /collabs", app.CreateCollab)
	http.HandleFunc("DELETE /collabs/{id}", app.DeleteCollab)
	http.HandleFunc("GET /collabs", app.GetAllCollabs)
	http.HandleFunc("GET /collabs/{id}", app.GetCollab)
	http.HandleFunc("PATCH /collabs/{id}", app.UpdateCollab)

	fmt.Println("API is ON at : http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	err := db.Conn.Ping()
	if err != nil {
		panic(err.Error())

	}
	fmt.Fprintf(w, "Connection réussi depuis %s", r.UserAgent())
}
