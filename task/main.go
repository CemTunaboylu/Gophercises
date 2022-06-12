package main

import (
	"task_manager/cmds"
	"task_manager/db"
	// homedir "github.com/mitchellh/go-homedir"
)

func main() {
	// home, _ := homedir.Dir()
	// dbPath := filepath.Join(home, "tasks.db")
	db.Connect(nil)
	cmds.Execute()

}
