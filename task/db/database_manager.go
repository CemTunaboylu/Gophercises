package db

type DBManager struct {
	host          string
	port          int
	user          string
	password      string
	database_name string
	table_name    string
}

func CreateDBManager(customizations ...func(dbm *DBManager)) *DBManager {
	d := &DBManager{
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

func (dbm *DBManager) SpitForConnection() []any {
	return []any{
		dbm.host,
		dbm.port,
		dbm.user,
		dbm.database_name,
	}
}

func (dbm *DBManager) SpitForInit() []any {
	return []any{
		dbm.host,
		dbm.port,
		dbm.user,
	}
}

func WithHost(host string) func(dbm *DBManager) {
	return func(d *DBManager) {
		d.host = host
	}
}

func WithPort(port int) func(dbm *DBManager) {
	return func(d *DBManager) {
		d.port = port
	}
}

func WithUser(user string) func(dbm *DBManager) {
	return func(d *DBManager) {
		d.user = user
	}
}

func WithDatabaseName(db_name string) func(dbm *DBManager) {
	return func(d *DBManager) {
		d.database_name = db_name
	}
}

func WithTableName(t_name string) func(dbm *DBManager) {
	return func(d *DBManager) {
		d.table_name = t_name
	}
}
