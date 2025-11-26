package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	driver   = "postgres"
	host     = "localhost"
	port     = "5432"
	user     = "diadieudoucoure"
	password = ""
	dbName   = "inventory_ntic"
)

var Conn *sql.DB

func NewDB() *sql.DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbName)
	conn, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database")
	Conn = conn
	return conn
}
