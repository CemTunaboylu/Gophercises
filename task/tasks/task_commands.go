package tasks

import (
	"log"
	"strconv"
	"task_manager/db"
)

// func Add manages the addition of the task t in the database - task list
func Add(t string) string {
	return db.Insert(t)
}

// func Do handles when the given task is done
func Do(ids ...string) string {
	ints := str_arr_to_any_arr(ids)
	return db.BulkMarkDone(ints)
}

// func List manages the printing of the current tasks to be done
func List(l int, v bool) (r string) {
	switch v {
	case true:
		r = db.VerboseListAll(l)
	case false:
		r = db.ListAll(l)
	}
	return
}

// func Remove updates the present boolean value of the tasks that are given to False, no real deletion
func Remove(all bool, ids ...string) string {
	ints := str_arr_to_any_arr(ids)
	return db.BulkRemove(all, ints)

}

func Completed() string {
	return db.Completed()
}

func str_arr_to_any_arr(s []string) []any {
	ints := make([]any, len(s))
	var err error
	for i, v := range s {
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			log.Printf("Given IDs are problematic %+v\n", s)
			return nil
		}
	}
	return ints

}
