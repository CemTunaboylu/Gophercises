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

var database *sql.DB
var db_manager *DBManager

func Init(dbm *DBManager) {
	if dbm == nil {
		dbm = CreateDBManager()
	}
	db_manager = dbm

	init_psql_str := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", dbm.SpitForInit()...)

	db, err := sql.Open("postgres", init_psql_str)
	Must(err)

	create_db := fmt.Sprintf("CREATE DATABASE %s", dbm.database_name)
	_, err = db.Exec(create_db)
	// Must(err)

	// _, err = db.Exec("USE " + dbm.database_name)
	// Must(err)

	database = db
	// create_database = `CREATE DATABASE task_manager OWNER cemtunaboylu;`

	table_cols := `
	id SERIAL PRIMARY KEY, 
	time_stamp TIMESTAMPTZ DEFAULT now(), 
	task VARCHAR(255), 
	done BOOLEAN DEFAULT False, 
	present BOOLEAN DEFAULT True, 
	last_modified TIMESTAMPTZ DEFAULT now()
	`
	init_cmd := fmt.Sprintf("CREATE TABLE %s(%s);", dbm.table_name, table_cols)
	_, err = database.Exec(init_cmd)
	Must(err)

	Connect(dbm)

}

func Connect(dbm *DBManager) {
	if dbm == nil {
		dbm = CreateDBManager()
	}
	db_manager = dbm
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%spassword=%s dbname=%s sslmode=disable",
	//  host, port, user, password, dbname)
	psql_str := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", dbm.SpitForConnection()...)

	db, err := sql.Open("postgres", psql_str)
	Must(err)
	err = db.Ping()
	Must(err)

	database = db
}

func Insert(t string) string {
	cmd := sql_insert_returning("id")
	assigned_id := 0
	// instead of Exec we use QueryRow indicating that we expect a row when returning
	// Exec normally have  Result has LastInsertId() (int64, error) and RowsAffected() (int64, error) methods, BUT lib/pq does not support this.
	err := database.QueryRow(cmd, t).Scan(&assigned_id)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	return fmt.Sprintf("(%d) '%s' is created\n", assigned_id, t)

}

func ListAll(num int) string {
	cmd := sql_select_w_lim(num, "done=False", "present=True")
	rows, err := database.Query(cmd)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	defer rows.Close()
	var tasks []string

	for rows.Next() {
		t := Task{}
		err = rows.Scan(&t.ID, &t.Time_Stamp, &t.Task_Text, &t.Is_Done, &t.Present, &t.Last_Modified)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, t.String())
	}
	return strings.Join(tasks, "")
}

func VerboseListAll(num int) string {
	cmd := sql_select_w_lim(num, "done=False", "present=True")
	rows, err := database.Query(cmd)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	defer rows.Close()
	var tasks []string

	for rows.Next() {
		t := Task{}
		err = rows.Scan(&t.ID, &t.Time_Stamp, &t.Task_Text, &t.Is_Done, &t.Present, &t.Last_Modified)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, t.V_String())
	}
	return strings.Join(tasks, "")
}

// Add time filter
func Completed() string {
	cmd := sql_select("done=True", "present=True")
	rows, err := database.Query(cmd)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}
	defer rows.Close()
	tasks := []string{"You have finished the following tasks today:\n"}
	for rows.Next() {
		t := Task{}
		err = rows.Scan(&t.ID, &t.Time_Stamp, &t.Task_Text, &t.Is_Done, &t.Present, &t.Last_Modified)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, t.C_String())
	}
	return strings.Join(tasks, "")
}

func BulkMarkDone(nums []any) string {
	var val_strs []string
	not_changed_map := map[any]struct{}{}

	for i, v := range nums {
		// create the appropriate value SQL strings in the form of $<num> for sql
		val_strs = append(val_strs, fmt.Sprintf(" $%d ", i+1))
		// our hashset for later identifying those who are not marked done
		not_changed_map[v] = struct{}{}
	}
	cmd := sql_update_returning("done=True", val_strs, "done=False", "present=True")
	rows, err := database.Query(cmd, nums...)
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}

	changed_id := 0
	var t string
	var results []string

	for rows.Next() {
		err = rows.Scan(&changed_id, &t)
		if err != nil {
			fmt.Println(err)
		}
		// removing the changed element
		delete(not_changed_map, changed_id)
		results = append(results, fmt.Sprintf("(%d) %s is marked as DONE \n", changed_id, t))
	}
	// handling messages of those who were already finished or is not there
	for k, _ := range not_changed_map {
		results = append(results, fmt.Sprintf("(%v) cannot be marked DONE\n", k))
	}
	return strings.Join(results, "")
}

func BulkRemove(all bool, nums []any) string {
	var val_strs []string
	not_changed_map := map[any]struct{}{}

	for i, v := range nums {
		// create the appropriate value SQL strings in the form of $<num> for sql
		val_strs = append(val_strs, fmt.Sprintf(" $%d ", i+1))
		// create our hashset for later identifying those who are not marked done
		not_changed_map[v] = struct{}{}
	}
	// cmd := `UPDATE ` + table_name + ` SET present=False, last_modified=now() WHERE id IN (` + strings.Join(val_strs, ",") + `) AND present=True RETURNING id, task`
	var cmd string
	var rows *sql.Rows
	var err error
	if all {
		cmd = sql_update_all_returning("present=False", "present=True")
		rows, err = database.Query(cmd)
	} else {
		cmd = sql_update_returning("present=False", val_strs, "present=True")
		rows, err = database.Query(cmd, nums...)
	}
	if err != nil {
		return fmt.Sprintf("%+v\n", err)
	}

	var changed_id int
	var t string
	var results []string

	for rows.Next() {
		err = rows.Scan(&changed_id, &t)
		if err != nil {
			fmt.Println(err)
		}
		// removing the changed element
		delete(not_changed_map, changed_id)
		results = append(results, fmt.Sprintf("(%d) %s is removed from the tasklist \n", changed_id, t))
	}
	// handling messages of those who were already finished or is not there
	for k, _ := range not_changed_map {
		results = append(results, fmt.Sprintf("(%v) cannot be removed from the tasklist\n", k))
	}
	return strings.Join(results, "")
}

func sql_select_w_lim(lim int, where ...string) string {
	return `SELECT * FROM ` + db_manager.table_name + ` WHERE ` + strings.Join(where, " AND ") + " ORDER BY time_stamp ASC LIMIT " + strconv.Itoa(lim)
}
func sql_select(where ...string) string {
	return `SELECT * FROM ` + db_manager.table_name + ` WHERE ` + strings.Join(where, " AND ") + ` ORDER BY time_stamp ASC`
}

func sql_insert_returning(ret ...string) string {
	return `INSERT INTO ` + db_manager.table_name + `(task) VALUES($1) RETURNING ` + strings.Join(ret, " , ")
}

func sql_update_returning(set string, in_val_strs []string, conds ...string) string {
	return `UPDATE ` + db_manager.table_name + ` SET last_modified=now(), ` + set +
		` WHERE id IN (` + strings.Join(in_val_strs, ",") + `) AND ` + strings.Join(conds, " AND ") + ` RETURNING id, task`

}

func sql_update_all_returning(set string, conds ...string) string {
	return `UPDATE ` + db_manager.table_name + ` SET last_modified=now(), ` + set +
		` WHERE ` + strings.Join(conds, " AND ") + ` RETURNING id, task`

}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
