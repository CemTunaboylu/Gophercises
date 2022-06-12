package db 

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	// The last import, _ "github.com/lib/pq", might look funny at first, but the short version is that we are importing the package so that it can register its drivers with the database/sql package,
	// and we use the _ identifier to tell Go that we still want this included even though we will never directly reference the package in our code.
	_ "github.com/lib/pq"
)


type PostgresManager struct {
	host          string
	port          int
	user          string
	password      string
	database_name string
	table_name    string
}
func CreatePostgresManager(customizations ...func(dbm *PostgresManager)) *PostgresManager {
	d := &PostgresManager{
		host:          "localhost",
		port:          5432,
		user:          "cemtunaboylu",
		password:      "",
		database_name: "task_manager",
		table_name:    "tasks",
	}
	for _, f := range customizations {
		f(d)
	}
	return d
}

func (dbm *PostgresManager) SpitForConnection() []any {
	return []any{
		dbm.host,
		dbm.port,
		dbm.user,
		dbm.database_name,
	}
}

func (dbm *PostgresManager) SpitForInit() []any {
	return []any{
		dbm.host,
		dbm.port,
		dbm.user,
	}
}

func WithHost(host string) func(dbm *PostgresManager) {
	return func(d *PostgresManager) {
		d.host = host
	}
}

func WithPort(port int) func(dbm *PostgresManager) {
	return func(d *PostgresManager) {
		d.port = port
	}
}

func WithUser(user string) func(dbm *PostgresManager) {
	return func(d *PostgresManager) {
		d.user = user
	}
}

func WithDatabaseName(db_name string) func(dbm *PostgresManager) {
	return func(d *PostgresManager) {
		d.database_name = db_name
	}
}

func WithTableName(t_name string) func(dbm *PostgresManager) {
	return func(d *PostgresManager) {
		d.table_name = t_name
	}
}
