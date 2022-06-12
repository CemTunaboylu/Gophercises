package db

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var db_name = "test_database"
var t_name = "test_tasks"
var test_db_manager = CreateDBManager(WithDatabaseName(db_name), WithTableName(t_name))

func TestCreateDataBase(t *testing.T) {
	// Init(test_db_manager)
	Connect(test_db_manager)
}

func TestInsert(t *testing.T) {
	Connect(test_db_manager)
	task := "Insertion Test"
	r := Insert(task)
	fmt.Println(r)

	if !strings.Contains(r, "created") {
		panic("TestInsert failed")
	}
}

func TestDone(t *testing.T) {
	task := "Done Test"
	Connect(test_db_manager)

	r := Insert(task)
	if !strings.Contains(r, "created") {
		panic("TestDone failed at Insertion")
	}
	fmt.Println(r)
	id := extract_id(r)
	r = BulkMarkDone([]any{id})
	fmt.Println(r)

	if !strings.Contains(r, "is marked as DONE") {
		panic("TestDone failed at Marking")
	}
}

func TestRemove(t *testing.T) {
	Connect(test_db_manager)

	task := "Remove Test"
	r := Insert(task)
	if !strings.Contains(r, "created") {
		panic("TestRemove failed at Insertion")
	}
	id := extract_id(r)
	r = BulkRemove(false, []any{id})
	if !strings.Contains(r, "is removed from") {
		panic("TestRemove failed at Removing")
	}
}

func TestRemoveAll(t *testing.T) {
	Connect(test_db_manager)

	task := "Remove All Test"
	r := Insert(task)
	if !strings.Contains(r, "created") {
		panic("TestRemoveAll failed at Insertion")
	}
	id := extract_id(r)
	r = BulkRemove(true, []any{id})
	fmt.Println(r)
	if !strings.Contains(r, "is removed from") {
		panic("TestRemoveAll failed at Removing")
	}
	VerboseListAll(10)
}

func extract_id(r string) int {
	id, err := strconv.Atoi(r[1:strings.Index(r, ")")])
	if err != nil {
		panic("Strconv failed.")
	}
	return id
}
