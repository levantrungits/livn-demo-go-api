// Package driver with Golang
// Designed by TRUNGLV
package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // lib Postgres Database
)

// Class PostgreDB data layer
type PostgresDB struct {
	SQL *sql.DB // pointer -> tranh tao ra 1 version copy
}

// Global varible
var Postgres = &PostgresDB{}

// Function Connect Postgres data base
func Connect(host, port, user, password, dbname string) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	Postgres.SQL = db
	return Postgres, nil
}
