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

func Connect(dbm *DBManager) *sql.DB {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%spassword=%s dbname=%s sslmode=disable",
	//  host, port, user, password, dbname)
	psql_str := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", dbm.SpitForConnection()...)

	db, err := sql.Open("postgres", psql_str)
	Must(err)
	err = db.Ping()
	Must(err)
	
	return db
}
func Insert(db *sql.DB, t string) string {
	assigned_id := 0
	// instead of Exec we use QueryRow indicating that we expect a row when returning
	// Exec normally have  Result has LastInsertId() (int64, error) and RowsAffected() (int64, error) methods, BUT lib/pq does not support this.
	err := database.QueryRow(cmd, t).Scan(&assigned_id)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return fmt.Sprintf("(%d) '%s' is created\n", assigned_id, t)

}







